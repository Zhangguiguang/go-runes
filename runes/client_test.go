package runes

import (
	"testing"
)

func TestXxx(t *testing.T) {
	c := NewClient(UrlTestnet)
	r, err := c.MintRunes()
	t.Log(r.Value, err)
}
