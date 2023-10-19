package types

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

const (
	// Define constants for your module's message types and routes
	TypeGenerateKeys = "register_keys"
)

func NewMsgRegisterKeys(account types.AccAddress, keys []string) *MsgRegisterKeys {
	return &MsgRegisterKeys{
		Account:   account,
		Keys:      keys,
		Timestamp: time.Now().Unix(),   // Current timestamp in Unix format or
		UniqueId:  uuid.New().String(), // A new UUID
	}
}

func (msg MsgRegisterKeys) Route() string { return RouterKey }

func (msg MsgRegisterKeys) Type() string { return TypeGenerateKeys }

func (msg MsgRegisterKeys) ValidateBasic() error {
	if len(msg.Account) == 0 {
		return sdkerrors.ErrInvalidAddress
	}
	if len(msg.Keys) == 0 {
		return fmt.Errorf("no keys to register")
	}
	return nil
}

func (msg MsgRegisterKeys) GetSignBytes() []byte {
	return types.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgRegisterKeys) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Account}
}
