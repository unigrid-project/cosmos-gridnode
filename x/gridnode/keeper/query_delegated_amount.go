package keeper

import (
	"context"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DelegatedAmount(goCtx context.Context, req *types.QueryDelegatedAmountRequest) (*types.QueryDelegatedAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Define the key for the delegated amount, assuming it's based on the delegator's address
	delegatorAddr, err := sdk.AccAddressFromBech32(req.DelegatorAddress)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid delegator address")
	}
	key := k.keyForDelegator(delegatorAddr)

	// Retrieve the value from the store
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(key)
	if bz == nil {
		// Return zero if no amount is found for the delegator
		return &types.QueryDelegatedAmountResponse{Amount: 0}, nil
	}

	// Convert the byte value to the appropriate data type
	amount := sdk.NewIntFromBigInt(new(big.Int).SetBytes(bz))

	return &types.QueryDelegatedAmountResponse{Amount: amount.Int64()}, nil
}
