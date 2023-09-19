package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetRegistrationFees = "set_registration_fees"

var _ sdk.Msg = &MsgSetRegistrationFees{}

func NewMsgSetRegistrationFees(creator string, domain string, feesByName string, feesByLength string, defaultFees string) *MsgSetRegistrationFees {
	return &MsgSetRegistrationFees{
		Creator:      creator,
		Domain:       domain,
		FeesByName:   feesByName,
		FeesByLength: feesByLength,
		DefaultFees:  defaultFees,
	}
}

func (msg *MsgSetRegistrationFees) Route() string {
	return RouterKey
}

func (msg *MsgSetRegistrationFees) Type() string {
	return TypeMsgSetRegistrationFees
}

func (msg *MsgSetRegistrationFees) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetRegistrationFees) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetRegistrationFees) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
