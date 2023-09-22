package keeper

import (
	"errors"
	"fmt"
	"math/big"

	sdkmath "cosmossdk.io/math"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
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
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, delegator, types.ModuleName, sdk.NewCoins(sdk.NewCoin("ugd", amount)))
	if err != nil {
		return err
	}

	// Store the locked tokens in the gridnode module's state
	lockedBalance := k.GetLockedBalance(ctx, delegator)
	lockedBalance = lockedBalance.Add(amount)
	k.SetLockedBalance(ctx, delegator, lockedBalance)

	return nil
}

// UndelegateTokens unlocks the tokens from gridnode delegation
func (k Keeper) UndelegateTokens(ctx sdk.Context, delegator sdk.AccAddress, amount sdkmath.Int) error {
	// ... similar logic to release the tokens
	fmt.Println("UndelegateTokens: ", delegator, amount)
	return nil
}

// Helper functions to get and set locked balance in the state
func (k Keeper) GetLockedBalance(ctx sdk.Context, delegator sdk.AccAddress) sdkmath.Int {
	// ... retrieve the locked balance from the store
	fmt.Println("GetLockedBalance: ", delegator)
	return sdk.NewInt(0)
}

func (k Keeper) SetLockedBalance(ctx sdk.Context, delegator sdk.AccAddress, amount sdkmath.Int) {
	// ... set the locked balance in the store
	fmt.Println("SetLockedBalance: ", delegator, amount)
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

func (k Keeper) CastVoteFromGridnode(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress, option govtypes.VoteOption) error {
	// Check if the voter is a Gridnode
	if !k.IsGridnode(ctx, voterAddr) {
		return errors.New("address is not a Gridnode")
	}

	// Use the x/gov keeper to add the vote
	_, err := k.govKeeper.AddVote(ctx, proposalID, voterAddr, option)
	return err
}

func (k Keeper) IsGridnode(ctx sdk.Context, voterAddr sdk.AccAddress) bool {
	res := true

	if res {
		return true
	}

	return false
}
