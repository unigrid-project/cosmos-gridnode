package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// Define constants for your module's message types and routes
	TypeMsgDelegate = "delegategridnode"
)

func NewMsgDelegateGridnode(delegatorAddress string, amount int64) *MsgGridnodeDelegate {
	fmt.Println("NewMsgDelegateGridnode: ", delegatorAddress, amount)
	return &MsgGridnodeDelegate{
		DelegatorAddress: delegatorAddress,
		Amount:           amount,
	}
}

func NewMsgUndelegateGridnode(delegatorAddress sdk.AccAddress, amount int64) *MsgGridnodeUndelegate {
	return &MsgGridnodeUndelegate{
		DelegatorAddress: delegatorAddress.String(), // Convert to string
		Amount:           amount,
	}
}

// Implementing the sdk.Msg interface

func (msg *MsgGridnodeDelegate) Route() string {
	fmt.Println("Route:", RouterKey)
	fmt.Println("msg:", msg)
	return RouterKey
}

func (msg *MsgGridnodeDelegate) Type() string {
	fmt.Println("TypeMsgDelegate:", TypeMsgDelegate)
	fmt.Println("msg:", msg)
	return TypeMsgDelegate
}

func (msg MsgGridnodeDelegate) ValidateBasic() error {
	fmt.Println("Validating MsgGridnodeDelegate:", msg)
	fmt.Println("Amount:", msg.Amount)
	fmt.Println("Delegator:", msg.DelegatorAddress)
	if msg.DelegatorAddress == "" {
		fmt.Println("Validation Failed: DelegatorAddress is empty")
		return sdkerrors.ErrInvalidAddress
	}
	if msg.Amount <= 0 {
		fmt.Println("Validation Failed: Amount is less than or equal to zero")
		return sdkerrors.ErrInvalidCoins
	}
	fmt.Println("Validation Successful")
	return nil
}

func (msg MsgGridnodeDelegate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgGridnodeDelegate) GetSigners() []sdk.AccAddress {
	delegatorAddr, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{delegatorAddr}
}

func (msg *MsgGridnodeUndelegate) Route() string {
	return RouterKey
}

func (msg *MsgGridnodeUndelegate) Type() string {
	return "undelegate"
}

func (msg MsgGridnodeUndelegate) ValidateBasic() error {
	fmt.Println("Delegator Address:", msg.DelegatorAddress)
	if msg.DelegatorAddress == "" {
		return sdkerrors.ErrInvalidAddress
	}
	if msg.Amount <= 0 {
		return sdkerrors.ErrInvalidCoins
	}
	return nil
}

func (msg MsgGridnodeUndelegate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgGridnodeUndelegate) GetSigners() []sdk.AccAddress {
	delegatorAddr, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{delegatorAddr}
}
