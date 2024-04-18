package rune

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestnetApi() *Api {
	return NewApi(UrlTestnet)
}

func TestTransferRunes(t *testing.T) {
	api := TestnetApi()

	sender := "tb1p37zl29ltg5x0nsmd8lf2m5fzrjealrrektcng03ztvmlcf3ejdqsh4tfht"
	receiver := "tb1p37zl29ltg5x0nsmd8lf2m5fzrjealrrektcng03ztvmlcf3ejdqsh4tfht"
	rune := "RUNE•TO•MOON"
	amount := "0.1"
	txn, err := api.TransferRunes(sender, receiver, rune, amount, 1.0, 500)
	require.NoError(t, err)
	t.Log(txn.JsonString())
}

func TestRuneBalance(t *testing.T) {
	api := TestnetApi()

	owner := "tb1p37zl29ltg5x0nsmd8lf2m5fzrjealrrektcng03ztvmlcf3ejdqsh4tfht"
	rune := "RUNE•TO•MOON"
	res, err := api.RuneBalance(owner, rune)
	require.NoError(t, err)
	t.Log(res)
	t.Log(res.BalanceWithDecimal())
}

func TestRuneBalances(t *testing.T) {
	api := TestnetApi()

	owner := "tb1p37zl29ltg5x0nsmd8lf2m5fzrjealrrektcng03ztvmlcf3ejdqsh4tfht"
	res, err := api.RuneBalances(owner)
	require.NoError(t, err)
	t.Log(res.JsonString())
}

func TestRuneInfo(t *testing.T) {
	api := TestnetApi()

	runeName := "HELLO•WORLD•FAIR"
	info, err := api.RuneInfo(runeName)
	require.NoError(t, err)
	t.Log(info.JsonString())
}

func TestRuneInfoById(t *testing.T) {
	api := TestnetApi()

	id := "2584614:143"
	info, err := api.RuneInfoById(id)
	require.NoError(t, err)
	t.Log(info.JsonString())
}
