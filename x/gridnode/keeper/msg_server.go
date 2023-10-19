package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/types"
)

type msgServer struct {
	types.UnimplementedMsgServer
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = &msgServer{}

func (s *msgServer) DelegateTokens(ctx context.Context, req *types.MsgGridnodeDelegate) (*types.MsgGridnodeDelegateResponse, error) {
	fmt.Println("msgServer DelegateTokens called with:", req)
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
		Status: "success",
	}, nil
}

func (s *msgServer) UndelegateTokens(ctx context.Context, req *types.MsgGridnodeUndelegate) (*types.MsgGridnodeUndelegateResponse, error) {
	// Convert context.Context to sdk.Context
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Extract information from the request
	delegatorAddr, err := sdk.AccAddressFromBech32(req.DelegatorAddress)
	if err != nil {
		return nil, err
	}
	amount := sdk.NewInt(req.Amount)

	// Call the Keeper's UndelegateTokens method to perform the business logic
	err = s.Keeper.UndelegateTokens(sdkCtx, delegatorAddr, amount)
	if err != nil {
		return nil, err
	}

	// Construct and return a response
	return &types.MsgGridnodeUndelegateResponse{
		Status: "success",
	}, nil
}

func (s *msgServer) RegisterKeys(ctx context.Context, req *types.MsgRegisterKeys) (*types.MsgRegisterKeysResponse, error) {
	// Convert context.Context to sdk.Context
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Convert []byte to sdk.AccAddress
	account := sdk.AccAddress(req.Account)

	// Call the Keeper's RegisterKeys method to perform the business logic
	err := s.Keeper.RegisterKeys(sdkCtx, account, req.Keys, req.Timestamp, req.UniqueId)
	if err != nil {
		return nil, err
	}

	// Construct and return a response
	return &types.MsgRegisterKeysResponse{
		Status: "success",
	}, nil
}
