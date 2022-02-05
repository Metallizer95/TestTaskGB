package usecases

type UseCases interface {
	FindMaxBalanceWalletForLastBlocks(numberBlocks int64) (WalletBalanceModel, error)
}
