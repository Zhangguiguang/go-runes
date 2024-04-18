package rune

import (
	"github.com/coming-chat/wallet-SDK/core/base"
	"github.com/coming-chat/wallet-SDK/core/btc"
)

type TransferTransaction struct {
	CommitId     string                 `json:"commit_id"`
	CommitPsbt   string                 `json:"commit_psbt"`
	NetworkFee   int64                  `json:"network_fee"`
	SatpointFee  int64                  `json:"satpoint_fee"`
	ServiceFee   int64                  `json:"service_fee"`
	CommitFee    int64                  `json:"commit_fee"`
	CommitVsize  int64                  `json:"commit_vsize"`
	CommitCustom *btc.Brc20CommitCustom `json:"commit_custom"`
}

// return networkFee + serviceFee + satpointFee
func (t *TransferTransaction) TotalFee() int64 {
	return t.NetworkFee + t.ServiceFee + t.SatpointFee
}

func (t *TransferTransaction) JsonString() (*base.OptionalString, error) {
	return base.JsonString(t)
}

// MARK - Implement base.Transaction

func (t *TransferTransaction) SignWithAccount(account base.Account) (signedTxn *base.OptionalString, err error) {
	return nil, base.ErrUnsupportedFunction
}

func (t *TransferTransaction) SignedTransactionWithAccount(account base.Account) (signedTxn base.SignedTransaction, err error) {
	defer base.CatchPanicAndMapToBasicError(&err)

	btcAccount := account.(*btc.Account)
	if btcAccount == nil {
		return nil, base.ErrInvalidAccountType
	}
	return btcAccount.SignPsbt(t.CommitPsbt)
}
