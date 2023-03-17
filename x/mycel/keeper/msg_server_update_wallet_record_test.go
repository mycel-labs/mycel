package keeper_test

import (
	"testing"

	"mycel/testutil"
	"mycel/x/mycel/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func GetMsgUpdateWalletRecord() *types.MsgUpdateWalletRecord {
	return &types.MsgUpdateWalletRecord{
		Creator:          testutil.Alice,
		Name:             "poyo",
		Parent:           "ninniku",
		WalletRecordType: "ETHEREUM_MAINNET",
		Value:            "0x1234567890123456789012345678901234567890",
	}
}

func TestUpdateWalletRecordSuccess(t *testing.T) {
	msgServer, keeper, context := setupMsgServer(t)
	// Register domain
	domain := GetMsgRegisterDomain()
	_, err := msgServer.RegisterDomain(context, domain)
	require.Nil(t, err)

	// Update wallet record
	record := GetMsgUpdateWalletRecord()
	_, err = msgServer.UpdateWalletRecord(context, record)
	require.Nil(t, err)

	// Check if wallet record is UpdateWalletRecord
	ctx := sdk.UnwrapSDKContext(context)
	res, _ := keeper.GetDomain(ctx, domain.Name, domain.Parent)
	require.Equal(t, record.Value, res.WalletRecords[record.WalletRecordType].Value)
}
