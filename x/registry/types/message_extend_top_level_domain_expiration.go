package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgExtendTopLevelDomainExpiration = "extend_top_level_domain_expiration"

var _ sdk.Msg = &MsgExtendTopLevelDomainExpiration{}

func NewMsgExtendTopLevelDomainExpiration(creator string, name string, registrationPeriodInYear int64) *MsgExtendTopLevelDomainExpiration {
	return &MsgExtendTopLevelDomainExpiration{
		Creator:                  creator,
		Name:                     name,
		RegistrationPeriodInYear: registrationPeriodInYear,
	}
}

func (msg *MsgExtendTopLevelDomainExpiration) Route() string {
	return RouterKey
}

func (msg *MsgExtendTopLevelDomainExpiration) Type() string {
	return TypeMsgExtendTopLevelDomainExpiration
}

func (msg *MsgExtendTopLevelDomainExpiration) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgExtendTopLevelDomainExpiration) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgExtendTopLevelDomainExpiration) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
