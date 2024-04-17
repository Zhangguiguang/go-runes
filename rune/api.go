package rune

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func NewRunesApi(url string) *Api {
	if !strings.HasSuffix(url, "/") {
		url = url + "/"
	}
	return &Api{Url: url}
}

func (c *Api) TransferRunes() (*base.OptionalString, error) {
	return base.NewOptionalString(""), nil
}

func (c *Api) RunesBalance(owner, rune string) (*Balance, error) {
	path := fmt.Sprintf("RuneBalance/%v/%v", owner, rune)
	var out Balance
	err := c.doRequest(http.MethodGet, path, nil, &out)
	return &out, err
}

func (c *Api) RunesBalances(owner string) (*BalanceArray, error) {
	path := fmt.Sprintf("RuneBalances/%v", owner)
	var out BalanceArray
	err := c.doRequest(http.MethodGet, path, nil, &out)
	return &out, err
}

func (c *Api) RuneInfo(name string) (*Info, error) {
	path := fmt.Sprintf("RuneInfo/%v", name)
	var out Info
	err := c.doRequest(http.MethodGet, path, nil, &out)
	return &out, err
}

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
