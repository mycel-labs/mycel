package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateDnsRecord = "update_dns_record"

var _ sdk.Msg = &MsgUpdateDnsRecord{}

func NewMsgUpdateDnsRecord(creator string, name string, parent string, dnsRecordType string, value string) *MsgUpdateDnsRecord {
	return &MsgUpdateDnsRecord{
		Creator:       creator,
		Name:          name,
		Parent:        parent,
		DnsRecordType: dnsRecordType,
		Value:         value,
	}
}

func (msg *MsgUpdateDnsRecord) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDnsRecord) Type() string {
	return TypeMsgUpdateDnsRecord
}

func (msg *MsgUpdateDnsRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDnsRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDnsRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
