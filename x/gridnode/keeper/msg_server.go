package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{Keeper: *keeper}
}

var _ types.MsgServer = msgServer{}

var _ types.GridnodeMsgServer = &msgServer{}

func (s *msgServer) DelegateTokens(ctx context.Context, req *types.MsgGridnodeDelegate) (*types.MsgGridnodeDelegateResponse, error) {
	// Convert context.Context to sdk.Context
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Extract information from the request
	delegatorAddr, err := sdk.AccAddressFromBech32(req.DelegatorAddress)
	if err != nil {
		return nil, err
	}
	amount := sdk.NewInt(req.Amount)

	// Call the Keeper's DelegateTokens method to perform the business logic
	err = s.Keeper.DelegateTokens(sdkCtx, delegatorAddr, amount)
	if err != nil {
		return nil, err
	}

	// Construct and return a response
	return &types.MsgGridnodeDelegateResponse{
		// Populate with any needed response data
	}, nil
}
