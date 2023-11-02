package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mycel-domain/mycel/app/params"
)

func NewBurnAmount(config EpochBurnConfig, index uint64) BurnAmount {
	return BurnAmount{
		Index:                 index,
		TotalEpochs:           config.DefaultTotalEpochs,
		CurrentEpoch:          0,
		TotalBurnAmount:       sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(0)),
		CumulativeBurntAmount: sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(0)),
	}
}

func (burnAmount BurnAmount) CalculateBurntAmount() sdk.Coin {
	if burnAmount.TotalBurnAmount.Amount.GTE(sdk.NewInt(int64(burnAmount.TotalEpochs))) {
		quotient := burnAmount.TotalBurnAmount.Amount.QuoRaw(int64(burnAmount.TotalEpochs))
		remander := burnAmount.TotalBurnAmount.Amount.ModRaw(int64(burnAmount.TotalEpochs))
		if remander.IsZero() || burnAmount.CurrentEpoch+1 != burnAmount.TotalEpochs {
			return sdk.NewCoin(params.DefaultBondDenom, quotient)
		}
		return sdk.NewCoin(params.BaseCoinUnit, quotient.Add(remander))
	} else if burnAmount.CurrentEpoch == 0 {
		return sdk.NewCoin(params.DefaultBondDenom, burnAmount.TotalBurnAmount.Amount)
	}
	return sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(0))
}

func (burnAmount *BurnAmount) CumulateBurntAmount(amount sdk.Coin) {
	burnAmount.CumulativeBurntAmount = burnAmount.CumulativeBurntAmount.Add(amount)
	burnAmount.CurrentEpoch++
}
