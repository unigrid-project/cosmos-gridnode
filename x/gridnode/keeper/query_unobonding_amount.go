package keeper

import (
	"context"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UnbondingEntries(goCtx context.Context, req *types.QueryUnbondingEntriesRequest) (*types.QueryUnbondingEntriesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Convert the delegator address string to sdk.AccAddress
	delegatorAddr, err := sdk.AccAddressFromBech32(req.BondingAccountAddress)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid delegator address")
	}

	// Define the key for the unbonding entries based on the delegator's address and block height
	key := k.keyForUnBonding(delegatorAddr, ctx.BlockHeight())

	// Retrieve the value from the store
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(key)
	if bz == nil {
		// Return an empty list if no unbonding entries are found for the delegator
		return &types.QueryUnbondingEntriesResponse{UnbondingEntries: nil}, nil
	}

	// Deserialize the byte value to a list of unbonding entries
	var entries []types.UnbondingEntry
	if err := json.Unmarshal(bz, &entries); err != nil {
		return nil, status.Error(codes.Internal, "failed to deserialize unbonding entries")
	}

	// Convert slice of UnbondingEntry to slice of pointers to UnbondingEntry
	entriesPtr := make([]*types.UnbondingEntry, len(entries))
	for i := range entries {
		entriesPtr[i] = &entries[i]
	}

	return &types.QueryUnbondingEntriesResponse{UnbondingEntries: entriesPtr}, nil
}
