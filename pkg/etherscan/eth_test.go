package etherscan

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLastTag(t *testing.T) {
	eth := New()
	lastTag, err := eth.GetLastBlockTag()
	assert.NoError(t, err)
	assert.NotEqual(t, "", lastTag)
	fmt.Println(lastTag)
}

func TestGetBlockByTag(t *testing.T) {
	eth := New()
	lastTag, err := eth.GetLastBlockTag()
	assert.NoError(t, err)

	block, err := eth.GetBlockByTag(lastTag)
	assert.NoError(t, err)
	assert.NotEqual(t, EthModel{}, block)

	fmt.Println(block.Result.Transactions)
}