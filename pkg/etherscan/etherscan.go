package etherscan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TODO: Create config struct and config file

type action string

const (
	BlockByTagAction   action = "eth_getBlockByNumber"
	LastBlockTagAction action = "eth_blockNumber"
)

type Etherscan interface {
	GetLastBlockTag() (string, error)
	GetBlockByTag(tag string) (EthModel, error)
}

type addressParameters struct {
	module  string
	action  action
	tag     string
	boolean string
	apikey  string
}

type etherscan struct {
	rootAddress string
	params      addressParameters
}

func New() Etherscan {
	return etherscan{
		rootAddress: "https://api.etherscan.io/api",
		params: addressParameters{
			module:  "proxy",
			action:  "",
			tag:     "",
			boolean: "true", // transaction details visible
			apikey:  "JPXGPM5DID3KFD3VH77TKI2RCU7Q3GMYPT",
		},
	}
}

func (e etherscan) GetLastBlockTag() (string, error) {
	e.params.action = LastBlockTagAction

	address := e.getRequestAddress()
	resp, err := http.Get(address)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	rawDataFromHttp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var responseModel LastTagModel
	err = json.Unmarshal(rawDataFromHttp, &responseModel)
	if err != nil {
		return "", err
	}

	return responseModel.Tag, nil
}

func (e etherscan) GetBlockByTag(tag string) (EthModel, error) {
	e.params.action = BlockByTagAction
	e.params.tag = tag

	address := e.getRequestAddress()
	fmt.Println(address)
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

func (e etherscan) getRequestAddress() string {
	var result string
	if e.params.action == BlockByTagAction {
		result = fmt.Sprintf(
			"%s?module=%s&action=%s&tag=%s&boolean=%s&apikey=%s",
			e.rootAddress,
			e.params.module,
			e.params.action,
			e.params.tag,
			e.params.boolean,
			e.params.apikey,
		)
	} else {
		result = fmt.Sprintf(
			"%s?module=%s&action=%s&apikey=%s",
			e.rootAddress,
			e.params.module,
			e.params.action,
			e.params.apikey,
		)
	}
	return result
}
