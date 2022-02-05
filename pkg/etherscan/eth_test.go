package etherscan

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Do not use api key in tests, so may occurred error related with limit of requests

func TestGetLastTag(t *testing.T) {
	eth := New(Config{
		RootAddress: "https://api.etherscan.io/api",
		Module:      "proxy",
	})
	lastTag, err := eth.GetLastBlockTag()
	assert.NoError(t, err)
	assert.NotEqual(t, 0, lastTag)
	fmt.Println(lastTag)
}

func TestGetBlockByTag(t *testing.T) {
	eth := New(Config{
		RootAddress: "https://api.etherscan.io/api",
		Module:      "proxy",
	})
	lastTag, err := eth.GetLastBlockTag()
	assert.NoError(t, err)

	block, err := eth.GetBlockByTag(lastTag)
	assert.NoError(t, err)
	assert.NotEqual(t, EthModel{}, block)

	fmt.Println(block.Result.Transactions)
}
