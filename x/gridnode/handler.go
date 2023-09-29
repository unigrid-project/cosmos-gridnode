package gridnode

// import (
// 	"fmt"
// 	"strconv"

// 	"cosmossdk.io/errors"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/keeper"
// 	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/types"
// )

// func NewHandler(am AppModule) sdk.Handler {
// 	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
// 		switch msg := msg.(type) {
// 		case *types.MsgGridnodeDelegate:
// 			fmt.Println("MsgGridnodeDelegate received in NewHandler: ", msg)
// 			return handleMsgDelegate(ctx, am, msg)
// 		case *types.MsgGridnodeUndelegate:
// 			return handleMsgUndelegate(ctx, am.keeper, msg)
// 		default:
// 			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
// 			return nil, errors.Wrapf(types.ErrUnknownRequest, errMsg)
// 		}

// 	}
// }

// func handleMsgDelegate(ctx sdk.Context, am AppModule, msg *types.MsgGridnodeDelegate) (*sdk.Result, error) {
// 	fmt.Println("handleMsgDelegate: ", msg)
// 	delegatorAddr, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Retrieve the available balance of the delegator account
// 	availableBalance := am.bankKeeper.GetBalance(ctx, delegatorAddr, "ugd")
// 	fmt.Println("availableBalance: ", availableBalance)
// 	// Retrieve the amount already delegated by the delegator
// 	delegatedAmount := am.keeper.GetDelegatedAmount(ctx, delegatorAddr)
// 	fmt.Println("delegatedAmount: ", delegatedAmount)
// 	// Calculate the maximum amount the delegator can delegate
// 	maxDelegatable := availableBalance.Amount.Sub(delegatedAmount)
// 	fmt.Println("maxDelegatable: ", maxDelegatable)

// 	msgAmount := sdk.NewInt(msg.Amount) // Convert int64 to sdk.Int
// 	// Check if the delegator has enough balance to delegate the specified amount
// 	if msgAmount.GT(maxDelegatable) {
// 		return nil, errors.Wrapf(types.ErrInsufficientFunds, "account %s has insufficient funds to delegate %s", delegatorAddr, strconv.FormatInt(msg.Amount, 10))
// 	}

// 	fmt.Println("handleMsgDelegate: ", msg)
// 	amount := sdk.NewInt(msg.Amount)
// 	err = am.keeper.DelegateTokens(ctx, delegatorAddr, amount)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Return a result or event to indicate successful delegation
// 	ctx.EventManager().EmitEvent(sdk.NewEvent(
// 		types.EventTypeDelegate,
// 		sdk.NewAttribute(types.AttributeKeyDelegator, msg.DelegatorAddress),
// 		sdk.NewAttribute(types.AttributeKeyAmount, strconv.FormatInt(msg.Amount, 10)),
// 	))
// 	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
// }

// func handleMsgUndelegate(ctx sdk.Context, k keeper.Keeper, msg *types.MsgGridnodeUndelegate) (*sdk.Result, error) {
// 	// Your logic for handling the MsgGridnodeUndelegate message goes here
// 	fmt.Println("handleMsgUndelegate: ", msg)
// 	return &sdk.Result{}, nil
// }
