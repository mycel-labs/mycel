package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawRegistrationFee = "withdraw_registration_fee"

var _ sdk.Msg = &MsgWithdrawRegistrationFee{}

func NewMsgWithdrawRegistrationFee(creator string, name string) *MsgWithdrawRegistrationFee {
	return &MsgWithdrawRegistrationFee{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgWithdrawRegistrationFee) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawRegistrationFee) Type() string {
	return TypeMsgWithdrawRegistrationFee
}

func (msg *MsgWithdrawRegistrationFee) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawRegistrationFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawRegistrationFee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
