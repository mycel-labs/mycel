package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateWalletRecord = "update_wallet_record"

var _ sdk.Msg = &MsgUpdateWalletRecord{}

func NewMsgUpdateWalletRecord(creator string, name string, parent string, walletRecordType string, value string) *MsgUpdateWalletRecord {
	return &MsgUpdateWalletRecord{
		Creator:          creator,
		Name:             name,
		Parent:           parent,
		WalletRecordType: walletRecordType,
		Value:            value,
	}
}

func (msg *MsgUpdateWalletRecord) Route() string {
	return RouterKey
}

func (msg *MsgUpdateWalletRecord) Type() string {
	return TypeMsgUpdateWalletRecord
}

func (msg *MsgUpdateWalletRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateWalletRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateWalletRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
