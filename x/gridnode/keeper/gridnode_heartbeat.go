package keeper

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	store "github.com/cosmos/cosmos-sdk/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/types"
)

const (
	heartbeatURL = "http://127.0.0.1:5000/gridnode/heartbeat"
	interval     = 1 * time.Minute
	hashKey      = "lastHashKey"
)

type HeartbeatManager struct {
	StoreKey store.StoreKey
	Keeper   *Keeper
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

func (hm *HeartbeatManager) SendHeartbeatIfDataChanged(ctx sdk.Context, data []Delegation) {
	fmt.Println("Checking for data changes...")
	store := ctx.KVStore(hm.StoreKey)
	newHashBytes := sha256.Sum256([]byte(fmt.Sprintf("%v", data)))
	newHash := hex.EncodeToString(newHashBytes[:])

	oldHashBytes := store.Get([]byte(hashKey))
	oldHash := string(oldHashBytes)

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
}

func (hm *HeartbeatManager) sendHeartbeat(data []Delegation) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(heartbeatURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
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

	for {
		select {
		case <-ticker.C:
			fmt.Println("Fetching delegation data...")
			data := hm.GetDelegationData(ctx)
			fmt.Printf("Fetched %d delegations.\n", len(data))
			fmt.Println("Sending heartbeat if data changed...")
			hm.SendHeartbeatIfDataChanged(ctx, data)
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

func (hm *HeartbeatManager) GetDelegationData(ctx sdk.Context) []Delegation {
	delegations, err := hm.Keeper.QueryAllDelegations(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
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

	return simplifiedDelegations
}
