package entity

type Asset struct {
	Name            string `yaml:"name"`
	Decimals        int    `yaml:"decimals"`
	Symbol          string `yaml:"symbol"`
	LogoPNG         string `yaml:"logo_png"`
	ContractAddress string `yaml:"contract_address"`
	Type            string `yaml:"type"`
	IsStableCoin    bool   `yaml:"is_stable_coin"`
	IsWrapper       bool   `yaml:"is_wrapper"`
	IsNFT           bool   `yaml:"is_nft"`
	IsDefi          bool   `yaml:"is_defi"`
}
