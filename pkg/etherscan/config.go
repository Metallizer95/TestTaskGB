package etherscan

type Config struct {
	RootAddress string `yaml:"root_address"`
	Module      string `yaml:"module"`
	ApiKey      string `yaml:"api_key"`
}
