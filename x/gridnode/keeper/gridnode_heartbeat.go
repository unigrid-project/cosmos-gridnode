package keeper

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"

	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/spf13/viper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/types"
)

const (
	interval = 1 * time.Minute
	hashKey  = "lastHashKey"
)

type HeartbeatManager struct {
	hedgehogUrl string
	StoreKey    store.StoreKey
	Keeper      *Keeper
}

type Delegation struct {
	Account         string `json:"account"`
	DelegatedAmount int64  `json:"delegated_amount"`
}

func NewHeartbeatManager(storeKey store.StoreKey, keeper *Keeper) *HeartbeatManager {
	return &HeartbeatManager{
		StoreKey: storeKey,
		Keeper:   keeper,
	}
}

func (hm *HeartbeatManager) SendHeartbeatIfDataChanged(ctx sdk.Context, data []Delegation) error {
	fmt.Println("Checking for data changes...")

	// Sort the data slice based on the Account field before hashing it
	sort.Slice(data, func(i, j int) bool {
		return data[i].Account < data[j].Account
	})

	store := ctx.KVStore(hm.StoreKey)
	newHashBytes := sha256.Sum256([]byte(fmt.Sprintf("%v", data)))
	newHash := hex.EncodeToString(newHashBytes[:])
	fmt.Printf("New Hash: %s", newHash)
	oldHashBytes := store.Get([]byte(hashKey))
	oldHash := string(oldHashBytes)
	fmt.Printf("Old Hash: %s", oldHash)
	if oldHash != newHash {
		fmt.Println("Data has changed, sending heartbeat...")
		err := hm.sendHeartbeat(data)
		if err != nil {
			// Handle error
			fmt.Println("Failed to send heartbeat:", err)
		} else {
			fmt.Println("Heartbeat sent successfully.")
			store.Set([]byte(hashKey), []byte(newHash))
		}
	} else {
		fmt.Println("No data change detected.")
	}
	return nil
}

func (hm *HeartbeatManager) sendHeartbeat(data []Delegation) error {
	base := viper.GetString("hedgehog.hedgehog_url")
	heartbeatURL := base + "/gridnode/heartbeat"
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("POST", heartbeatURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send heartbeat: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send heartbeat, status code: %d", resp.StatusCode)
	}

	return nil
}

func (hm *HeartbeatManager) StartHeartbeatTimer(ctx sdk.Context) {
	fmt.Println("Initializing heartbeat timer...")
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Check for a cancellation signal to stop the goroutine
	ctxDone := ctx.Done()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Fetching delegation data...")
			data, err := hm.GetDelegationData(ctx)
			if err != nil {
				fmt.Printf("Error fetching delegations: %v\n", err)
				continue // Skip this iteration if there's an error
			}
			fmt.Printf("Fetched %d delegations.\n", len(data))
			fmt.Println("Sending heartbeat if data changed...")
			err = hm.SendHeartbeatIfDataChanged(ctx, data)
			if err != nil {
				fmt.Printf("Error sending heartbeat: %v\n", err)
			}
		case <-ctxDone:
			// Context was cancelled, exit the goroutine
			fmt.Println("Heartbeat timer cancelled, exiting...")
			return
		}
	}
}

func processDelegations(chunk []types.DelegationInfo, results chan<- Delegation, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, delegation := range chunk {
		simplified := Delegation{
			Account:         delegation.Account,
			DelegatedAmount: delegation.DelegatedAmount,
		}
		results <- simplified
	}
}

func (hm *HeartbeatManager) GetDelegationData(ctx sdk.Context) ([]Delegation, error) {
	delegations, err := hm.Keeper.QueryAllDelegations(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err // Return the error along with a nil slice
	}

	numWorkers := 4 // adjust this based on your needs
	numDelegations := len(delegations)
	chunkSize := numDelegations / numWorkers
	results := make(chan Delegation, numDelegations)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numWorkers-1 {
			end = numDelegations // ensure any remaining items are processed
		}
		wg.Add(1)
		go processDelegations(delegations[start:end], results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	simplifiedDelegations := []Delegation{}
	for result := range results {
		simplifiedDelegations = append(simplifiedDelegations, result)
	}

	return simplifiedDelegations, nil
}

func (hm *HeartbeatManager) SetHedgehogUrl(url string) {
	hm.hedgehogUrl = url
}

func (hm *HeartbeatManager) GetHedgehogUrl() string {
	return hm.hedgehogUrl
}
