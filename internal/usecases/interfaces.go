package usecases

type UseCases interface {
	FindMaxBalanceWalletForLastBlocks(numberBlocks int) (WalletBalanceModel, error)
}
