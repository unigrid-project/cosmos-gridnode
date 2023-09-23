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

// type MsgGridnodeDelegate struct {
// 	DelegatorAddress sdk.AccAddress `json:"delegator_address" yaml:"delegator_address"`
// 	Amount           int64          `json:"amount" yaml:"amount"`
// }

// type MsgGridnodeUndelegate struct {
// 	DelegatorAddress sdk.AccAddress `json:"delegator_address" yaml:"delegator_address"`
// 	Amount           int64          `json:"amount" yaml:"amount"`
// }

// ProtoMessage implements types.Msg.
// func (MsgGridnodeDelegate) ProtoMessage() {
// 	panic("unimplemented")
// }

// // Reset implements types.Msg.
// func (MsgGridnodeDelegate) Reset() {
// 	panic("unimplemented")
// }

// // String implements types.Msg.
// func (MsgGridnodeDelegate) String() string {
// 	panic("unimplemented")
// }

// func (MsgGridnodeUndelegate) ProtoMessage() {
// 	panic("unimplemented")
// }

// // Reset implements types.Msg.
// func (MsgGridnodeUndelegate) Reset() {
// 	panic("unimplemented")
// }

// // String implements types.Msg.
// func (MsgGridnodeUndelegate) String() string {
// 	panic("unimplemented")
// }

func NewMsgDelegateGridnode(delegatorAddress sdk.AccAddress, amount int64) *MsgGridnodeDelegate {
	fmt.Println("NewMsgDelegateGridnode: ", delegatorAddress, amount)
	return &MsgGridnodeDelegate{
		DelegatorAddress: delegatorAddress.String(), // Convert to string
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
	fmt.Println("Delegator Address:", msg.DelegatorAddress)
	if msg.DelegatorAddress == "" {
		return sdkerrors.ErrInvalidAddress
	}
	if msg.Amount <= 0 {
		return sdkerrors.ErrInvalidCoins
	}
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
