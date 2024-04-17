package btc

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

type RunesApi struct {
	Url string
}

func NewRunesApi(url string) *RunesApi {
	if !strings.HasSuffix(url, "/") {
		url = url + "/"
	}
	return &RunesApi{Url: url}
}

func (c *RunesApi) TransferRunes() (*base.OptionalString, error) {
	return base.NewOptionalString(""), nil
}

func (c *RunesApi) RunesBalance(owner, rune string) (*RuneBalance, error) {
	path := fmt.Sprintf("RuneBalance/%v/%v", owner, rune)
	var out RuneBalance
	err := c.doRequest(http.MethodGet, path, nil, &out)
	return &out, err
}

func (c *RunesApi) RunesBalances(owner string) (*RuneBalanceArray, error) {
	path := fmt.Sprintf("RuneBalances/%v", owner)
	var out RuneBalanceArray
	err := c.doRequest(http.MethodGet, path, nil, &out)
	return &out, err
}

func (c *RunesApi) RuneInfo(name string) (*RuneInfo, error) {
	path := fmt.Sprintf("RuneInfo/%v", name)
	var out RuneInfo
	err := c.doRequest(http.MethodGet, path, nil, &out)
	return &out, err
}

func (c *RunesApi) RuneInfoById(id string) (*RuneInfo, error) {
	path := fmt.Sprintf("RuneInfoById/%v", id)
	var out RuneInfo
	err := c.doRequest(http.MethodGet, path, nil, &out)
	return &out, err
}

func (c *RunesApi) doRequest(method string, path string, params map[string]any, out any) (err error) {
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
