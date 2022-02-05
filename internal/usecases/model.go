package usecases

type WalletsModel struct {
	Wallets []WalletBalanceModel
}

type WalletBalanceModel struct {
	Address string
	Value   float64
}

type ErrorModel struct {
	Errs []error
}

func (em *ErrorModel) Append(err error) {
	em.Errs = append(em.Errs, err)
}
