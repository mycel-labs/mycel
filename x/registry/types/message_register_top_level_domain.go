package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterTopLevelDomain = "register_top_level_domain"

var _ sdk.Msg = &MsgRegisterTopLevelDomain{}

func NewMsgRegisterTopLevelDomain(creator string, name string, registrationPeriodInYear uint64) *MsgRegisterTopLevelDomain {
	return &MsgRegisterTopLevelDomain{
		Creator:                  creator,
		Name:                     name,
		RegistrationPeriodInYear: registrationPeriodInYear,
	}
}

func (msg *MsgRegisterTopLevelDomain) Route() string {
	return RouterKey
}

func (msg *MsgRegisterTopLevelDomain) Type() string {
	return TypeMsgRegisterTopLevelDomain
}

func (msg *MsgRegisterTopLevelDomain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterTopLevelDomain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterTopLevelDomain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
