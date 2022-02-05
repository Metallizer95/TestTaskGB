package usecases

type WalletsModel struct {
	Wallets []WalletBalanceModel
}

type WalletBalanceModel struct {
	Address string
	Value   float64
}

type ErrModel struct {
	Msg string
}
