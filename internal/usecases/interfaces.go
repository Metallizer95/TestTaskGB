package usecases

type UseCases interface {
	FindMaxProfitWalletForLastBlocks(numberBlocks int64) (WalletBalanceModel, ErrorModel)
}
