package etherscan

type EthModel struct {
	Id     int    `json:"id"`
	Result Result `json:"result"`
}

type Result struct {
	Transactions []TransactionModel `json:"transactions"`
}

type TransactionModel struct {
	SenderAddress   string `json:"from"`
	ReceiverAddress string `json:"to"`
	Value           string `json:"value"`
}

type LastTagModel struct {
	Tag string `json:"result"`
}
