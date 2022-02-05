package etherscan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// TODO: Create config struct and config file

const (
	BlockByTagAction   string = "eth_getBlockByNumber"
	LastBlockTagAction string = "eth_blockNumber"

	NumberFreeRPS int64         = 5
	RequestsDelay time.Duration = 300 * time.Millisecond
)

type Etherscan interface {
	GetLastBlockTag() (int64, error)
	GetBlockByTag(tag int64) (EthModel, error)
}

type etherScan struct {
	rootAddress string
	apiKey      string
	module      string
}

func New(cfg Config) Etherscan {
	return etherScan{
		rootAddress: cfg.RootAddress,
		apiKey:      cfg.ApiKey,
		module:      cfg.Module,
	}
}

func (e etherScan) GetLastBlockTag() (int64, error) {
	address := fmt.Sprintf("%s?module=%s&action=%s&apikey=%s",
		e.rootAddress,
		e.module,
		LastBlockTagAction,
		e.apiKey,
	)
	resp, err := http.Get(address)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	rawDataFromHttp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var responseModel LastTagModel
	err = json.Unmarshal(rawDataFromHttp, &responseModel)
	if err != nil {
		return 0, err
	}

	tagInt, err := strconv.ParseInt(responseModel.Tag[2:], 16, 64)
	if err != nil {
		return 0, err
	}

	return tagInt, nil
}

func (e etherScan) GetBlockByTag(tag int64) (EthModel, error) {
	hexTag := "0x" + strconv.FormatInt(tag, 16)

	address := fmt.Sprintf("%s?module=%s&action=%s&tag=%s&boolean=true&apikey=%s",
		e.rootAddress,
		e.module,
		BlockByTagAction,
		hexTag,
		e.apiKey,
	)
	resp, err := http.Get(address)
	if err != nil {
		return EthModel{}, err
	}
	defer resp.Body.Close()

	rawDataFromHttp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return EthModel{}, err
	}

	var responseModel EthModel
	err = json.Unmarshal(rawDataFromHttp, &responseModel)
	if err != nil {
		return EthModel{}, err
	}

	return responseModel, nil
}
