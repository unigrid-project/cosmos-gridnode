package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

const (
	// Define constants for your module's message types and routes
	TypeMsgDelegate   = "delegategridnode"
	TypeMsgUndelegate = "undelegategridnode"
)

func NewMsgDelegateGridnode(delegatorAddress string, amount int64) *MsgGridnodeDelegate {
	fmt.Println("NewMsgDelegateGridnode: ", delegatorAddress, amount)
	return &MsgGridnodeDelegate{
		DelegatorAddress: delegatorAddress,
		Amount:           amount,
		Timestamp:        time.Now().Unix(),   // Current timestamp in Unix format or
		UniqueId:         uuid.New().String(), // A new UUID
	}
}

func NewMsgUndelegateGridnode(delegatorAddress string, amount int64) *MsgGridnodeUndelegate {
	fmt.Println("NewMsgUndelegateGridnode: ", delegatorAddress, amount)
	return &MsgGridnodeUndelegate{
		DelegatorAddress: delegatorAddress,
		Amount:           amount,
		Timestamp:        time.Now().Unix(),   // Current timestamp in Unix format or
		UniqueId:         uuid.New().String(), // A new UUID
	}
}

// Implementing the sdk.Msg interface
func (msg *MsgGridnodeDelegate) Route() string {
	return RouterKey
}

func (msg *MsgGridnodeDelegate) Type() string {
	return TypeMsgDelegate
}

func (msg MsgGridnodeDelegate) ValidateBasic() error {
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
	return TypeMsgUndelegate
}

func (msg MsgGridnodeUndelegate) ValidateBasic() error {
	//fmt.Println("Delegator Address:", msg.DelegatorAddress)
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
