package usecases

import (
	"github.com/Metallizer95/TestTaskGB/internal/domain"
	"github.com/Metallizer95/TestTaskGB/pkg/etherscan"
	"github.com/cheggaaa/pb/v3"
	"log"
	"math/big"
	"sort"
	"time"
)

type useCasesImpl struct {
	EthScanService etherscan.Etherscan
}

func New(eth etherscan.Etherscan) UseCases {
	return useCasesImpl{
		EthScanService: eth,
	}
}

func (uc useCasesImpl) FindMaxProfitWalletForLastBlocks(numberBlocks int64) (wb WalletBalanceModel, errs ErrorModel) {
	lastBlock, err := uc.EthScanService.GetLastBlockTag()
	if err != nil {
		errs.Append(err)
		return wb, errs
	}

	holders := domain.NewHolders()

	// Do not use goroutines because RPS is restricted (only 5 per second for free version)
	bar := pb.StartNew(int(numberBlocks))
	defer bar.Finish()
	for i := lastBlock; i > lastBlock-numberBlocks; i-- {
		if i%etherscan.NumberFreeRPS == 0 && i != 0 {
			time.Sleep(etherscan.RequestsDelay)
		}
		iBlock, err := uc.EthScanService.GetBlockByTag(i)
		if err != nil {
			log.Printf("%d-iteration failed: %v", i, err)
			continue
		}

		for _, t := range iBlock.Result.Transactions {
			valBigFloat, _, _ := big.ParseFloat(t.Value[2:], 16, 0, 0)
			val, _ := valBigFloat.Float64()
			holders.AddValue(t.ReceiverAddress, val)
			holders.SubValue(t.SenderAddress, -val)
		}
		bar.Increment()
	}
	return uc.getWalletMaxProfit(holders), errs
}

func (uc useCasesImpl) getWalletMaxProfit(h domain.Holders) WalletBalanceModel {
	var wbs []WalletBalanceModel
	for k, v := range h.Holders {
		wbs = append(wbs, WalletBalanceModel{
			Address: k,
			Value:   v,
		})
	}
	sort.Slice(wbs, func(i, j int) bool {
		return wbs[i].Value > wbs[j].Value
	})
	return wbs[0]
}
