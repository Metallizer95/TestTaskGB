package etherscan

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var cfg Config

func TestMain(t *testing.M) {
	fullPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	sepPath := strings.Split(fullPath, "/")
	rootPath := strings.Join(sepPath[:len(sepPath)-2], "/")

	configPath := rootPath + "/ethConfig.yml"

	cfg, err = NewConfig(configPath)
	if err != nil {
		panic(err)
	}

	t.Run()
}

func TestGetLastTag(t *testing.T) {
	eth := New(cfg)
	lastTag, err := eth.GetLastBlockTag()
	assert.NoError(t, err)
	assert.NotEqual(t, 0, lastTag)
}

func TestGetBlockByTag(t *testing.T) {
	eth := New(cfg)
	lastTag, err := eth.GetLastBlockTag()
	assert.NoError(t, err)
	block, err := eth.GetBlockByTag(lastTag)
	assert.NoError(t, err)
	assert.NotEqual(t, EthModel{}, block)
}
