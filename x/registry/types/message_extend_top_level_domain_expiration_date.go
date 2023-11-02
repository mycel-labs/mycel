package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgExtendTopLevelDomainExpirationDate = "extend_top_level_domain_expiration"

var _ sdk.Msg = &MsgExtendTopLevelDomainExpirationDate{}

func NewMsgExtendTopLevelDomainExpirationDate(creator string, name string, extentsionPeriodInYear uint64) *MsgExtendTopLevelDomainExpirationDate {
	return &MsgExtendTopLevelDomainExpirationDate{
		Creator:                  creator,
		Name:                     name,
		ExtensionPeriodInYear	: extentsionPeriodInYear,
	}
}

func (msg *MsgExtendTopLevelDomainExpirationDate) Route() string {
	return RouterKey
}

func (msg *MsgExtendTopLevelDomainExpirationDate) Type() string {
	return TypeMsgExtendTopLevelDomainExpirationDate
}

func (msg *MsgExtendTopLevelDomainExpirationDate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgExtendTopLevelDomainExpirationDate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgExtendTopLevelDomainExpirationDate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
