package runes

import "github.com/coming-chat/wallet-SDK/core/base"

const (
	UrlMainnet = "https://ord.bevm.io/mainnet"
	UrlTestnet = "https://ord.bevm.io/testnet"
)

type Client struct {
	Url string
}

func NewClient(url string) *Client {
	return &Client{Url: url}
}

func (c *Client) MintRunes() (*base.OptionalString, error) {
	return base.NewOptionalString("sjiso"), nil
}
