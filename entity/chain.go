package entity

type ChainType string

const (
	MainChainType ChainType = "main"
	TestChainType           = "test"
)

type Chain struct {
	Id       string    `yaml:"id"`
	ParentId string    `yaml:"parent_id"`
	Name     string    `yaml:"name"`
	Symbol   string    `yaml:"symbol"`
	LogoPNG  string    `yaml:"logo_png"`
	ChainId  string    `yaml:"chain_id"`
	Assets   []*Asset  `yaml:"assets"`
	Type     ChainType `yaml:"type"`
}
