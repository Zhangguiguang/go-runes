package rune

import (
	"encoding/json"
	"math/big"

	"github.com/coming-chat/wallet-SDK/core/base"
	"github.com/coming-chat/wallet-SDK/core/base/inter"
)

// MARK - Balance

// rune balance
type Balance struct {
	Address        string `json:"addr"`
	Balance        string `json:"balance"`
	Divisibility   int16  `json:"divisibility"` // decimal
	HasInscription bool   `json:"has_inscription"`
	Rune           string `json:"rune"`
	Symbol         string `json:"symbol"`
}

func (b *Balance) UnmarshalJSON(data []byte) error {
	var raw struct {
		Address        string   `json:"addr"`
		Balance        *big.Int `json:"balance"` // only balance need use big int
		Divisibility   int16    `json:"divisibility"`
		HasInscription bool     `json:"has_inscription"`
		Rune           string   `json:"rune"`
		Symbol         string   `json:"symbol"`
	}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	b.Address = raw.Address
	b.Balance = raw.Balance.String()
	b.Divisibility = raw.Divisibility
	b.HasInscription = raw.HasInscription
	b.Rune = raw.Rune
	b.Symbol = raw.Symbol
	return nil
}

func (b *Balance) BalanceWithDecimal() string {
	balFloat, ok := big.NewFloat(0).SetString(b.Balance)
	if !ok {
		return "0"
	}
	pow := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(b.Divisibility)), nil)
	bal := big.NewFloat(0).Quo(balFloat, big.NewFloat(0).SetInt(pow))
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

// rune info
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
