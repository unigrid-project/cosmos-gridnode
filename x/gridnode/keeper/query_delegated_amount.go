package keeper

import (
	"context"
	"encoding/json"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DelegatedAmount(goCtx context.Context, req *types.QueryDelegatedAmountRequest) (*types.QueryDelegatedAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	delegatorAddr, err := sdk.AccAddressFromBech32(req.DelegatorAddress)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid delegator address")
	}
	key := k.keyForDelegator(delegatorAddr)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.StoreKey))
	bz := store.Get(key)
	if bz == nil {
		return &types.QueryDelegatedAmountResponse{Amount: 0}, nil
	}

	var delegationData types.DelegationData
	err = json.Unmarshal(bz, &delegationData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmarshal delegation data: %v", err)
	}

	return &types.QueryDelegatedAmountResponse{Amount: delegationData.LockedBalance.Int64()}, nil
}
