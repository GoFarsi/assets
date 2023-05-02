package entity

type Chain struct {
	Name    string   `yaml:"name"`
	Symbol  string   `yaml:"symbol"`
	LogoPNG string   `yaml:"logo_png"`
	ChainId string   `yaml:"chain_id"`
	Assets  []*Asset `yaml:"assets"`
}
