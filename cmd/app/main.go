package main

import (
	"flag"
	"fmt"
	"github.com/Metallizer95/TestTaskGB/internal/usecases"
	"github.com/Metallizer95/TestTaskGB/pkg/etherscan"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var ethConfigPath string

func init() {
	flag.StringVar(&ethConfigPath, "path", "./ethConfig.yml", "")
	flag.Parse()
}

func main() {
	cfgFile, err := ioutil.ReadFile(ethConfigPath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var ethCfg etherscan.Config
	err = yaml.Unmarshal(cfgFile, &ethCfg)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	ucs := usecases.New(etherscan.New(ethCfg))
	res, _ := ucs.FindMaxBalanceWalletForLastBlocks(20)
	fmt.Println(res)
}
