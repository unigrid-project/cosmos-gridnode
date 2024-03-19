package keeper

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"cosmossdk.io/core/store"
	"cosmossdk.io/errors"
	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/types"
)

type (
	Keeper struct {
		cdc              codec.BinaryCodec
		storeKey         storetypes.StoreKey
		memKey           storetypes.StoreKey
		paramstore       paramtypes.Subspace
		bankKeeper       types.BankKeeper
		govKeeper        types.GovKeeper
		heartbeatMgr     *HeartbeatManager
		heartbeatStarted bool
		authority        string
		storeService     store.KVStoreService
		accountKeeper    types.AccountKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	ps paramtypes.Subspace,
	bk types.BankKeeper,
	authority string,
	storeService store.KVStoreService,
	ak types.AccountKeeper,

) *Keeper {
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	keeper := &Keeper{
		cdc:           cdc,
		paramstore:    ps,
		bankKeeper:    bk,
		authority:     authority,
		storeService:  storeService,
		accountKeeper: ak,
	}

	//keeper.heartbeatMgr = NewHeartbeatManager(storeKey, keeper)
	// Initialize HeartbeatManager with storeService instead of storeKey
	keeper.heartbeatMgr = NewHeartbeatManager(storeService, keeper)
	return keeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// DelegateTokens locks the tokens for gridnode delegation
func (k Keeper) DelegateTokens(ctx context.Context, delegator sdk.AccAddress, amount sdkmath.Int) error {
	// Retrieve the available balance of the delegator account
	availableBalance := k.bankKeeper.GetBalance(ctx, delegator, "uugd")
	fmt.Println("availableBalance: ", availableBalance)
	// Retrieve the amount already delegated by the delegator
	delegatedAmount := k.GetDelegatedAmount(ctx, delegator)
	fmt.Println("delegatedAmount: ", delegatedAmount)
	// Calculate the maximum amount the delegator can delegate
	maxDelegatable := availableBalance.Amount.Sub(delegatedAmount)
	fmt.Println("maxDelegatable: ", maxDelegatable)

	// Retrieve the account using the delegator's address
	account := k.accountKeeper.GetAccount(ctx, delegator)
	if account == nil {
		return fmt.Errorf("account %s not found", delegator)
	}

	// Get the public key from the account
	pubKey := account.GetPubKey()
	if pubKey == nil {
		return fmt.Errorf("public key not found for account %s", delegator)
	}
	fmt.Printf("Public Key for delegator: %s\n", pubKey)
	// Convert the public key to a usable format or log it
	pubKeyBytes := pubKey.Bytes()
	fmt.Println("Public Key bytes for delegator:", pubKeyBytes)
	pubKeyHex := hex.EncodeToString(pubKeyBytes)
	fmt.Println("Public Key in Hex:", pubKeyHex)
	//pubKeyHex := PublicKeyToHex(pubKey)
	// Check if the delegator has enough balance to delegate the specified amount
	if amount.GT(maxDelegatable) {
		return errors.Wrapf(types.ErrInsufficientFunds, "account %s has insufficient funds to delegate %s uugd", delegator, amount.String())
	}

	// Deduct tokens from user's balance
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, delegator, types.ModuleName, sdk.NewCoins(sdk.NewCoin("uugd", amount)))
	if err != nil {
		return errors.Wrapf(types.ErrInsufficientFunds, "failed to delegate tokens: %v", err)
	}

	// Store the locked tokens in the gridnode module's state
	lockedBalance := k.GetLockedBalance(ctx, delegator)
	fmt.Println("Current Locked balance before adding: ", lockedBalance) // Log the current locked balance before adding the new amount
	lockedBalance = lockedBalance.Add(amount)
	fmt.Println("Locked balance after adding: ", lockedBalance) // Log the locked balance after adding the new amount
	k.SetLockedBalance(ctx, delegator, lockedBalance, pubKeyHex)
	uwCtx := sdk.UnwrapSDKContext(ctx)
	// Emitting events
	uwCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDelegate,
		sdk.NewAttribute(types.AttributeKeyDelegator, delegator.String()),
		sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
	))

	return nil
}

// UndelegateTokens unlocks the tokens from gridnode delegation
func (k Keeper) UndelegateTokens(ctx context.Context, account sdk.AccAddress, amount sdkmath.Int) error {
	// ... similar logic to release the tokens
	fmt.Println("UndelegateTokens: ", account, amount)
	// Retrieve the current unbonding entries for the account
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.StoreKey))

	key := k.keyForUnBonding(account)
	var currentUnbondingEntries []types.UnbondingEntry
	if bz := store.Get(key); bz != nil {
		if err := json.Unmarshal(bz, &currentUnbondingEntries); err != nil {
			return err
		}
	}

	// Calculate the total unbonding amount including the new amount
	totalUnbonding := amount
	for _, entry := range currentUnbondingEntries {
		totalUnbonding = totalUnbonding.Add(sdkmath.NewInt(entry.Amount))
	}

	// Retrieve the delegated amount for the account
	delegatedAmount := k.GetDelegatedAmount(ctx, account)

	// Check if the total unbonding amount exceeds the delegated amount
	if totalUnbonding.GT(delegatedAmount) {
		return errors.Wrapf(types.ErrOverUnbond, "attempting to unbond more than the delegated amount")
	}
	uwCtx := sdk.UnwrapSDKContext(ctx)
	// Retrieve current block time
	blockTime := uwCtx.BlockTime()

	// Define the unbonding period, 21 days TODO: enable this for mainnet
	//unbondingPeriod := time.Hour * 24 * 21
	// Define the unbonding period, (for testnet 1 day)
	//unbondingPeriod := time.Hour * 24 * 1
	// 2 minutes for testing
	unbondingPeriod := time.Minute * 2
	// Calculate the completion time for the unbonding
	completionTime := blockTime.Add(unbondingPeriod)

	// Create an UnbondingEntry
	entry := types.UnbondingEntry{
		Account:        account.String(),
		Amount:         amount.Int64(),
		CompletionTime: completionTime.Unix(),
	}

	// Store the unbonding entry in the state using the AddUnbonding method
	if err := k.AddUnbondingEntry(ctx, entry); err != nil {
		return err
	}

	// Emit an event or log the unbonding
	uwCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUnbond,
			sdk.NewAttribute(types.AttributeKeyDelegator, account.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
			sdk.NewAttribute(types.AttributeKeyCompletionTime, completionTime.String()),
		),
	)

	return nil
}

func (k Keeper) GetStoreKey() storetypes.StoreKey {
	return k.storeKey
}

func (k Keeper) GetBankKeeper() types.BankKeeper {
	return k.bankKeeper
}

func (k Keeper) GetStoreService() store.KVStoreService {
	return k.storeService
}

func (k Keeper) GetLockedBalance(ctx context.Context, delegator sdk.AccAddress) sdkmath.Int {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.StoreKey))
	key := k.keyForDelegator(delegator)

	fmt.Println("Getting Locked Balance for delegator: ", delegator, " with key: ", string(key))
	bz := store.Get(key)
	if bz == nil {
		fmt.Println("No Locked Balance found for delegator: ", delegator)
		return sdkmath.ZeroInt()
	}

	var delegationData types.DelegationData
	err := json.Unmarshal(bz, &delegationData)
	if err != nil {
		panic(err) // Handle the error appropriately. You might want to return an error instead of panicking.
	}

	fmt.Println("Found Locked Balance: ", delegationData.LockedBalance, " for delegator: ", delegator)
	return delegationData.LockedBalance
}

func (k Keeper) QueryAllDelegations(ctx context.Context) ([]types.DelegationInfo, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.StoreKey))

	delegatedAmountPrefixStore := prefix.NewStore(store, []byte(delegatedAmountPrefix))

	var delegations []types.DelegationInfo
	var mu sync.Mutex
	var wg sync.WaitGroup

	iterator := delegatedAmountPrefixStore.Iterator(nil, nil)

	if iterator == nil {
		return nil, errors.New("iterator is nil", 0, "QueryAllDelegations")
	}

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		wg.Add(1) // Increment the WaitGroup counter
		go func(key, value []byte) {
			defer wg.Done() // Decrement the counter when the goroutine completes

			// Ensure the key is long enough to slice
			if len(key) < len(delegatedAmountPrefix) {
				fmt.Printf("Key is too short: %s\n", key)
				return // or return an error
			}

			// Extract the account address directly
			accountAddr := string(key)

			// Deserialize the byte value to DelegationData
			var delegationData types.DelegationData
			err := json.Unmarshal(value, &delegationData)
			if err != nil {
				fmt.Printf("Error unmarshalling delegation data: %v. Raw data: %x\n", err, value)
				return // or return an error
			}

			// Convert the account address string to sdk.AccAddress
			delegatorAddr, err := sdk.AccAddressFromBech32(accountAddr)
			if err != nil {
				fmt.Printf("Error converting account address: %v\n", err)
				return // or return an error
			}

			// Define the key for the unbonding entries based on the delegator's address and block height
			unbondingKey := k.keyForUnBonding(delegatorAddr)

			// Retrieve the value from the store
			bz := store.Get(unbondingKey)
			if bz == nil {
				// If bz is nil, append a DelegationInfo object with an empty UnbondingEntries field
				info := types.DelegationInfo{
					Account:          accountAddr,
					DelegatedAmount:  delegationData.LockedBalance.Int64(),
					UnbondingEntries: nil, // UnbondingEntries is nil
					PublicKey:        delegationData.PublicKey,
				}
				mu.Lock() // Lock the mutex to prevent concurrent write to the slice
				delegations = append(delegations, info)
				mu.Unlock() // Unlock the mutex after writing to the slice
			} else {
				// Deserialize the byte value to a list of unbonding entries
				var unbondingEntries []types.UnbondingEntry
				err = json.Unmarshal(bz, &unbondingEntries)
				if err != nil {
					fmt.Printf("Error unmarshalling unbonding entries: %v\n", err)
					return // or return an error
				}
				// Convert slice of UnbondingEntry to slice of pointers to UnbondingEntry
				unbondingEntriesPtr := make([]*types.UnbondingEntry, len(unbondingEntries))
				for i := range unbondingEntries {
					unbondingEntriesPtr[i] = &unbondingEntries[i]
				}
				simpleUnbondingEntries := make([]*types.SimpleUnbondingEntry, len(unbondingEntriesPtr))
				for i, entry := range unbondingEntriesPtr {
					simpleUnbondingEntries[i] = &types.SimpleUnbondingEntry{
						Amount:         entry.Amount,
						CompletionTime: entry.CompletionTime,
					}
				}

				// Append a DelegationInfo object with the UnbondingEntries field populated
				info := types.DelegationInfo{
					Account:          accountAddr,
					DelegatedAmount:  delegationData.LockedBalance.Int64(),
					UnbondingEntries: simpleUnbondingEntries, // UnbondingEntries is populated
					PublicKey:        delegationData.PublicKey,
				}
				mu.Lock() // Lock the mutex to prevent concurrent write to the slice
				delegations = append(delegations, info)
				mu.Unlock() // Unlock the mutex after writing to the slice
			}

		}(iterator.Key(), iterator.Value()) // Pass key and value as arguments to the goroutine
	}

	wg.Wait() // Wait for all goroutines to complete

	return delegations, nil
}

func (k Keeper) SetLockedBalance(ctx context.Context, delegator sdk.AccAddress, amount sdkmath.Int, pubKey string) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.StoreKey))
	key := k.keyForDelegator(delegator)

	delegationData := types.DelegationData{
		LockedBalance: amount,
		PublicKey:     pubKey,
	}

	data, err := json.Marshal(delegationData)
	if err != nil {
		panic(err) // Handle the error appropriately
	}
	fmt.Println("Setting Locked Balance: ", amount, " for delegator: ", delegator, " with key: ", string(key), "pubKey: ", pubKey) // Log the amount being set, the delegator, and the key being used

	store.Set(key, data)
}

const delegatedAmountPrefix = "delegatedAmount-"

func (k Keeper) keyForDelegator(delegator sdk.AccAddress) []byte {
	return []byte(delegatedAmountPrefix + delegator.String())
}

const bondingPrefix = "bonding-"

func (k Keeper) GetBondingPrefix() string {
	return bondingPrefix
}

func (k Keeper) keyForUnBonding(delegator sdk.AccAddress) []byte {
	return []byte(fmt.Sprintf("%s%s", bondingPrefix, delegator.String()))
}

func (k Keeper) GetDelegatedAmount(ctx context.Context, delegator sdk.AccAddress) sdkmath.Int {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.StoreKey))
	byteValue := store.Get(k.keyForDelegator(delegator))
	if byteValue == nil {
		fmt.Println("No delegation found for address:", delegator)
		return sdkmath.ZeroInt()
	}

	var delegationData types.DelegationData
	err := json.Unmarshal(byteValue, &delegationData)
	if err != nil {
		panic(err) // Handle the error appropriately
	}

	fmt.Println("Delegated amount for address", delegator, "is:", delegationData.LockedBalance.Int64())
	return delegationData.LockedBalance
}

func (k Keeper) GetPublicKeyForDelegator(ctx context.Context, delegator sdk.AccAddress) (string, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.StoreKey))

	bz := store.Get(k.keyForDelegator(delegator))
	if bz == nil {
		return "", fmt.Errorf("delegator not found")
	}

	var delegationData types.DelegationData
	err := json.Unmarshal(bz, &delegationData)
	if err != nil {
		return "", err
	}

	return delegationData.PublicKey, nil
}

func (k Keeper) SetDelegatedAmount(ctx context.Context, delegator sdk.AccAddress, amount sdkmath.Int) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.StoreKey))
	if amount.IsNegative() {
		fmt.Println("Warning: Trying to set a negative delegation amount for address:", delegator)
		// Handle negative amounts, perhaps log an error or panic
	}
	store.Set(k.keyForDelegator(delegator), amount.BigInt().Bytes())
	fmt.Println("Set delegated amount for address", delegator, "to:", amount)
}

// AddUnbondingEntry adds a new unbonding entry for a given account.
func (k Keeper) AddUnbondingEntry(ctx context.Context, entry types.UnbondingEntry) error {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.StoreKey))
	delegatorAddr, err := sdk.AccAddressFromBech32(entry.Account)
	if err != nil {
		return err
	}
	key := k.keyForUnBonding(delegatorAddr)

	var entries []types.UnbondingEntry
	if bz := store.Get(key); bz != nil {
		// Deserialize the existing entries
		if err := json.Unmarshal(bz, &entries); err != nil {
			return err
		}
	}

	// Add the new entry to the list
	entries = append(entries, entry)

	// Serialize the updated list of entries
	bz, err := json.Marshal(entries)
	if err != nil {
		return err
	}
	store.Set(key, bz)

	return nil
}

func (k *Keeper) StartHeartbeatTimer(ctx sdk.Context) {
	if k.heartbeatMgr.started {
		fmt.Println("Heartbeat timer already started")
		return
	}
	fmt.Println("Starting the heartbeat timer")
	go k.heartbeatMgr.StartHeartbeatTimer(ctx)
}

func PublicKeyToHex(pubKey crypto.PubKey) string {
	return hex.EncodeToString(pubKey.Bytes())
}
