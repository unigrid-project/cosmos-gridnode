package keeper

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		bankKeeper types.BankKeeper
		govKeeper  types.GovKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	bk types.BankKeeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		bankKeeper: bk,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// DelegateTokens locks the tokens for gridnode delegation
func (k Keeper) DelegateTokens(ctx sdk.Context, delegator sdk.AccAddress, amount sdkmath.Int) error {
	// Retrieve the available balance of the delegator account
	availableBalance := k.bankKeeper.GetBalance(ctx, delegator, "ugd")
	fmt.Println("availableBalance: ", availableBalance)
	// Retrieve the amount already delegated by the delegator
	delegatedAmount := k.GetDelegatedAmount(ctx, delegator)
	fmt.Println("delegatedAmount: ", delegatedAmount)
	// Calculate the maximum amount the delegator can delegate
	maxDelegatable := availableBalance.Amount.Sub(delegatedAmount)
	fmt.Println("maxDelegatable: ", maxDelegatable)

	// Check if the delegator has enough balance to delegate the specified amount
	if amount.GT(maxDelegatable) {
		return errors.Wrapf(types.ErrInsufficientFunds, "account %s has insufficient funds to delegate %s", delegator, amount.String())
	}

	// Deduct tokens from user's balance
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, delegator, types.ModuleName, sdk.NewCoins(sdk.NewCoin("ugd", amount)))
	if err != nil {
		return errors.Wrapf(types.ErrInsufficientFunds, "failed to delegate tokens: %v", err)
	}

	// Store the locked tokens in the gridnode module's state
	lockedBalance := k.GetLockedBalance(ctx, delegator)
	fmt.Println("Current Locked balance before adding: ", lockedBalance) // Log the current locked balance before adding the new amount
	lockedBalance = lockedBalance.Add(amount)
	fmt.Println("Locked balance after adding: ", lockedBalance) // Log the locked balance after adding the new amount
	k.SetLockedBalance(ctx, delegator, lockedBalance)

	// Emitting events
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDelegate,
		sdk.NewAttribute(types.AttributeKeyDelegator, delegator.String()),
		sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
	))

	return nil
}

// UndelegateTokens unlocks the tokens from gridnode delegation
func (k Keeper) UndelegateTokens(ctx sdk.Context, account sdk.AccAddress, amount sdkmath.Int) error {
	// ... similar logic to release the tokens
	fmt.Println("UndelegateTokens: ", account, amount)
	// Retrieve the current unbonding entries for the account
	store := ctx.KVStore(k.storeKey)
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

	// Retrieve current block time
	blockTime := ctx.BlockTime()

	// Define the unbonding period, 21 days
	//unbondingPeriod := time.Hour * 24 * 21
	// Define the unbonding period, 1 minute (for testing)
	unbondingPeriod := time.Minute
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
	ctx.EventManager().EmitEvent(
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

func (k Keeper) GetLockedBalance(ctx sdk.Context, delegator sdk.AccAddress) sdkmath.Int {
	store := ctx.KVStore(k.storeKey)
	key := k.keyForDelegator(delegator)

	fmt.Println("Getting Locked Balance for delegator: ", delegator, " with key: ", string(key)) // Log the delegator and the key being used to get the balance
	bz := store.Get(key)
	if bz == nil {
		fmt.Println("No Locked Balance found for delegator: ", delegator) // Log if no balance is found for the delegator
		return sdkmath.ZeroInt()
	}
	amount := sdkmath.NewIntFromBigInt(new(big.Int).SetBytes(bz))
	fmt.Println("Found Locked Balance: ", amount, " for delegator: ", delegator) // Log the amount found for the delegator
	return amount
}

func (k Keeper) QueryAllDelegations(ctx sdk.Context) ([]types.DelegationInfo, error) {
	store := ctx.KVStore(k.storeKey)

	if store == nil {
		return nil, errors.New("store is nil", 0, "QueryAllDelegations")
	}

	delegatedAmountPrefixStore := prefix.NewStore(store, []byte(delegatedAmountPrefix))

	var delegations []types.DelegationInfo

	iterator := delegatedAmountPrefixStore.Iterator(nil, nil)

	if iterator == nil {
		return nil, errors.New("iterator is nil", 0, "QueryAllDelegations")
	}

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		value := iterator.Value()
		fmt.Printf("Key: %s, Value: %x\n", key, value)

		if len(key) < len(delegatedAmountPrefix) {
			fmt.Printf("Key is too short: %s\n", key)
			continue // or return an error
		}

		// Parse the delegator address from the key
		delegatorAddr := sdk.AccAddress(key[len(delegatedAmountPrefix):])

		// Parse the delegated amount from the value
		delegatedAmount := sdkmath.NewIntFromBigInt(new(big.Int).SetBytes(value))

		fmt.Printf("Delegator Address: %s, Delegated Amount: %s\n", delegatorAddr, delegatedAmount)

		// Get unbonding entries for the account
		unbondingKey := k.keyForUnBonding(delegatorAddr)
		var unbondingEntries []types.UnbondingEntry = []types.UnbondingEntry{}

		if bz := store.Get(unbondingKey); bz != nil {
			if err := json.Unmarshal(bz, &unbondingEntries); err != nil {
				fmt.Printf("Error unmarshalling unbonding entries: %v\n", err)
				return nil, err
			}

		}

		// Sum up the unbonding amounts
		var unbondingAmount sdkmath.Int
		for _, entry := range unbondingEntries {
			unbondingAmount = unbondingAmount.Add(sdkmath.NewInt(entry.Amount))
		}

		fmt.Printf("Unbonding Entries: %v, Unbonding Amount: %s\n", unbondingEntries, unbondingAmount)

		info := types.DelegationInfo{
			Account:         delegatorAddr.String(),
			DelegatedAmount: delegatedAmount.Int64(),
			UnbondingAmount: unbondingAmount.Int64(),
		}
		delegations = append(delegations, info)
	}

	return delegations, nil
}

func (k Keeper) SetLockedBalance(ctx sdk.Context, delegator sdk.AccAddress, amount sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	key := k.keyForDelegator(delegator)
	fmt.Println("Setting Locked Balance: ", amount, " for delegator: ", delegator, " with key: ", string(key)) // Log the amount being set, the delegator, and the key being used
	store.Set(key, amount.BigInt().Bytes())
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

func (k Keeper) GetDelegatedAmount(ctx sdk.Context, delegator sdk.AccAddress) sdkmath.Int {
	store := ctx.KVStore(k.storeKey)
	byteValue := store.Get(k.keyForDelegator(delegator))
	if byteValue == nil {
		fmt.Println("No delegation found for address:", delegator)
		return sdkmath.ZeroInt()
	}
	amount := sdkmath.NewIntFromBigInt(new(big.Int).SetBytes(byteValue))
	fmt.Println("Delegated amount for address", delegator, "is:", amount)
	return amount
}

func (k Keeper) SetDelegatedAmount(ctx sdk.Context, delegator sdk.AccAddress, amount sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	if amount.IsNegative() {
		fmt.Println("Warning: Trying to set a negative delegation amount for address:", delegator)
		// Handle negative amounts, perhaps log an error or panic
	}
	store.Set(k.keyForDelegator(delegator), amount.BigInt().Bytes())
	fmt.Println("Set delegated amount for address", delegator, "to:", amount)
}

// AddUnbondingEntry adds a new unbonding entry for a given account.
func (k Keeper) AddUnbondingEntry(ctx sdk.Context, entry types.UnbondingEntry) error {
	store := ctx.KVStore(k.storeKey)
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
