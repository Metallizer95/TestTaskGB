package main

import (
	"flag"
	"fmt"
	"github.com/Metallizer95/TestTaskGB/internal/usecases"
	"github.com/Metallizer95/TestTaskGB/pkg/etherscan"
	"log"
)

var ethConfigPath string

func init() {
	flag.StringVar(&ethConfigPath, "path", "./ethConfig.yml", "")
	flag.Parse()
}

func main() {
	ethCfg, err := etherscan.NewConfig(ethConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	ucs := usecases.New(etherscan.New(ethCfg))
	res, errs := ucs.FindMaxProfitWalletForLastBlocks(100)

	fmt.Printf("Top holder: %s\nHis profit: %e wei", res.Address, res.Value)
	if len(errs.Errs) != 0 {
		fmt.Printf("Errors: %+v", errs)
	}
}
