package keeper

import (
	"fmt"
	"math/big"

	sdkmath "cosmossdk.io/math"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
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
	// Deduct tokens from user's balance
	fmt.Println("DelegateTokens: ", delegator, amount)
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, delegator, types.ModuleName, sdk.NewCoins(sdk.NewCoin("ugd", amount)))
	if err != nil {
		return err
	}

	// Store the locked tokens in the gridnode module's state
	lockedBalance := k.GetLockedBalance(ctx, delegator)
	fmt.Println("Current Locked balance before adding: ", lockedBalance) // Log the current locked balance before adding the new amount
	lockedBalance = lockedBalance.Add(amount)
	fmt.Println("Locked balance after adding: ", lockedBalance) // Log the locked balance after adding the new amount
	k.SetLockedBalance(ctx, delegator, lockedBalance)

	return nil
}

// UndelegateTokens unlocks the tokens from gridnode delegation
func (k Keeper) UndelegateTokens(ctx sdk.Context, delegator sdk.AccAddress, amount sdkmath.Int) error {
	// ... similar logic to release the tokens
	fmt.Println("UndelegateTokens: ", delegator, amount)
	return nil
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
