package usecases

import "github.com/Metallizer95/TestTaskGB/pkg/etherscan"

type useCasesImpl struct {
	EthScanService etherscan.Etherscan
}

func New() UseCases {
	return nil
}

func (uc *useCasesImpl) FindMaxBalanceWalletForLastBlocks(numberBlocks int) (WalletBalanceModel, error) {
	return WalletBalanceModel{}, nil
}
