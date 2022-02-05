package usecases

import (
	"fmt"
	"github.com/Metallizer95/TestTaskGB/internal/domain"
	"github.com/Metallizer95/TestTaskGB/pkg/etherscan"
	"math/big"
	"sort"
	"time"
)

type useCasesImpl struct {
	EthScanService etherscan.Etherscan
}

func New() UseCases {
	return useCasesImpl{
		EthScanService: etherscan.New(),
	}
}

func (uc useCasesImpl) FindMaxBalanceWalletForLastBlocks(numberBlocks int64) (WalletBalanceModel, error) {
	lastBlock, err := uc.EthScanService.GetLastBlockTag()
	if err != nil {
		return WalletBalanceModel{}, err
	}

	holders := domain.NewHolders()

	// Do not use goroutines because RPS is restricted (only 5 per second for free version)
	for i := lastBlock; i > lastBlock-numberBlocks; i-- {
		if i%5 == 0 && i != 0 {
			time.Sleep(300 * time.Millisecond)
		}
		iBlock, err := uc.EthScanService.GetBlockByTag(i)
		if err != nil {
			fmt.Println(err)
		}

		for _, t := range iBlock.Result.Transactions {
			valBigFloat, _, _ := big.ParseFloat(t.Value[2:], 16, 0, 0)
			val, _ := valBigFloat.Float64()
			holders.AddValue(t.ReceiverAddress, val)
			holders.SubValue(t.SenderAddress, -val)
		}
	}

	var ss []WalletBalanceModel
	for k, v := range holders.Holders {
		ss = append(ss, WalletBalanceModel{
			Address: k,
			Value:   v,
		})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	return ss[0], nil
}
