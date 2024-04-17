package rune

import (
	"math/big"

	"github.com/coming-chat/wallet-SDK/core/base"
	"github.com/coming-chat/wallet-SDK/core/base/inter"
)

// MARK - Balance

type Balance struct {
	Address        string `json:"addr"`
	Balance        int64  `json:"balance"`
	Divisibility   int16  `json:"divisibility"` // decimal
	HasInscription bool   `json:"has_inscription"`
	Rune           string `json:"rune"`
	Symbol         string `json:"symbol"`
}

func (b *Balance) BalanceWithDecimal() string {
	pow := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(b.Divisibility)), nil)
	bal := big.NewFloat(0).Quo(big.NewFloat(float64(b.Balance)), big.NewFloat(0).SetInt(pow))
	return bal.String()
}

func (b *Balance) JsonString() string {
	j, err := base.JsonString(b)
	if err != nil {
		return ""
	}
	return j.Value
}

type BalanceArray struct {
	inter.AnyArray[*Balance]
}

// MARK - Info

type Info struct {
	RuneId       string  `json:"rune_id"`
	Block        int64   `json:"block"`
	Burned       int64   `json:"burned"`
	Divisibility int16   `json:"divisibility"`
	Etching      string  `json:"etching"`
	Mints        int64   `json:"mints"`
	Number       int64   `json:"number"`
	Premine      int64   `json:"premine"`
	SpacedRune   string  `json:"spaced_rune"`
	Symbol       string  `json:"symbol"`
	Timestamp    int64   `json:"timestamp"`
	Progress     float64 `json:"progress"`
	//  Terms string`json:"terms"`
}

func (i *Info) JsonString() string {
	j, err := base.JsonString(i)
	if err != nil {
		return ""
	}
	return j.Value
}
