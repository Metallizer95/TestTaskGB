package etherscan

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	RootAddress string `yaml:"root_address"`
	Module      string `yaml:"module"`
	ApiKey      string `yaml:"api_key"`
}

func NewConfig(filepath string) (Config, error) {
	cfgFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return Config{}, err
	}

	var ethCfg Config
	err = yaml.Unmarshal(cfgFile, &ethCfg)
	if err != nil {
		return Config{}, err
	}

	return ethCfg, nil
}
