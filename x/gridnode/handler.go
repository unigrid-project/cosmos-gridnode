package gridnode

import (
	"fmt"
	"strconv"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/keeper"
	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/types"
)

func NewHandler(am AppModule) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.MsgDelegate:
			return handleMsgDelegate(ctx, am, msg)
		case types.MsgUndelegate:
			return handleMsgUndelegate(ctx, am.keeper, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, errors.Wrapf(types.ErrUnknownRequest, errMsg)

		}
	}
}

func handleMsgDelegate(ctx sdk.Context, am AppModule, msg types.MsgDelegate) (*sdk.Result, error) {
	coins := am.bankKeeper.SpendableCoins(ctx, msg.DelegatorAddress)
	// Check if the delegator has enough coins to delegate
	if coins.AmountOf("ugd").LT(sdk.NewInt(msg.Amount)) {
		return nil, errors.Wrapf(types.ErrInsufficientFunds, "not enough coins to delegate")
	}

	fmt.Println("handleMsgDelegate: ", msg)
	amount := sdk.NewInt(msg.Amount)
	err := am.keeper.DelegateTokens(ctx, msg.DelegatorAddress, amount)
	if err != nil {
		return nil, err
	}

	// Return a result or event to indicate successful delegation
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDelegate,
		sdk.NewAttribute(types.AttributeKeyDelegator, msg.DelegatorAddress.String()),
		sdk.NewAttribute(types.AttributeKeyAmount, strconv.FormatInt(msg.Amount, 10)),
	))
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUndelegate(ctx sdk.Context, k keeper.Keeper, msg types.MsgUndelegate) (*sdk.Result, error) {
	// Your logic for handling the MsgUndelegate message goes here
	fmt.Println("handleMsgUndelegate: ", msg)
	return &sdk.Result{}, nil
}
