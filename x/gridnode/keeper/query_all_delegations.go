package keeper

import (
	// ... other imports
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/types"
)

// AllDelegations implements the QueryServer.AllDelegations method
func (k Keeper) AllDelegations(ctx context.Context, req *types.QueryAllDelegationsRequest) (*types.QueryAllDelegationsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	delegations, err := k.QueryAllDelegations(sdkCtx)
	if err != nil {
		return nil, err
	}

	// Convert each DelegationInfo value to a pointer
	delegationPointers := make([]*types.DelegationInfo, len(delegations))
	for i, delegation := range delegations {
		delegationCopy := delegation // Optional: Copy the delegation if necessary
		delegationPointers[i] = &delegationCopy
	}

	return &types.QueryAllDelegationsResponse{
		Delegations: delegationPointers,
		Pagination:  &query.PageResponse{},
	}, nil
}
