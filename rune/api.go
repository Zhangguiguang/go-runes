package rune

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/coming-chat/wallet-SDK/core/base"
	"github.com/coming-chat/wallet-SDK/pkg/httpUtil"
)

const (
	UrlMainnet = "https://ord.bevm.io/mainnet"
	UrlTestnet = "https://ord.bevm.io/testnet"
)

type Api struct {
	Url string
}

func NewApi(url string) *Api {
	if !strings.HasSuffix(url, "/") {
		url = url + "/"
	}
	return &Api{Url: url}
}

// 构建 rune 转账交易
// @param postage: the rune utxo's satoshi value, default is 546
func (c *Api) TransferRunes(sender, receiver string, runeName, amount string, feeRate float64, postage int) (*TransferTransaction, error) {
	path := "TransferRunes"
	if postage <= 0 {
		postage = 546
	}
	params := map[string]any{
		"commit_fee_rate": strconv.FormatFloat(feeRate, 'f', -1, 64),
		"source":          sender,
		"receive_infos": []map[string]string{
			{
				"destination": receiver,
				"rune":        runeName,
				"amount":      amount,
			}},
		"postage": postage,
	}
	var out TransferTransaction
	err := c.doRequest(http.MethodPost, path, params, &out)
	return &out, err
}

// 查询一个地址下某个 rune 的 balance
func (c *Api) RuneBalance(owner, runeName string) (*Balance, error) {
	path := fmt.Sprintf("RuneBalance/%v/%v", owner, runeName)
	var out Balance
	err := c.doRequest(http.MethodGet, path, nil, &out)
	return &out, err
}

// 查询一个地址下所有的 rune balance
func (c *Api) RuneBalances(owner string) (*BalanceArray, error) {
	path := fmt.Sprintf("RuneBalances/%v", owner)
	var out BalanceArray
	err := c.doRequest(http.MethodGet, path, nil, &out)
	return &out, err
}

// 使用 rune 名称查询 rune 信息
func (c *Api) RuneInfo(name string) (*Info, error) {
	path := fmt.Sprintf("RuneInfo/%v", name)
	var out Info
	err := c.doRequest(http.MethodGet, path, nil, &out)
	return &out, err
}

// 使用 rune id 查询 rune 信息
func (c *Api) RuneInfoById(id string) (*Info, error) {
	path := fmt.Sprintf("RuneInfoById/%v", id)
	var out Info
	err := c.doRequest(http.MethodGet, path, nil, &out)
	return &out, err
}

func (c *Api) doRequest(method string, path string, params map[string]any, out any) (err error) {
	defer base.CatchPanicAndMapToBasicError(&err)

	url := c.Url + strings.TrimPrefix(path, "/")
	header := map[string]string{"Content-Type": "application/json"}

	paramsBytes, err := json.Marshal(params)
	if err != nil {
		return
	}
	resp, err := httpUtil.Request(method, url, header, paramsBytes)
	if err != nil {
		return
	}
	if resp.Code != http.StatusOK {
		return fmt.Errorf("code: %d, body: %s", resp.Code, string(resp.Body))
	}

	return json.Unmarshal(resp.Body, out)
}
