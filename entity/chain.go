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

func NewChain(name, symbol string, chainType ChainType, chainId string, assets ...*Asset) *Chain {
	return &Chain{
		Name:    name,
		Symbol:  symbol,
		ChainId: chainId,
		Type:    chainType,
		Assets:  assets,
	}
}

func (c *Chain) GetName() string {
	return c.Name
}

func (c *Chain) GetSymbol() string {
	return c.Symbol
}

func (c *Chain) GetChainId() string {
	return c.ChainId
}

func (c *Chain) GetChainType() ChainType {
	return c.Type
}

func (c *Chain) IsMainNet() bool {
	if c.Type == MainChainType {
		return true
	}

	return false
}

func (c *Chain) IsTestNet() bool {
	if c.Type == TestChainType {
		return true
	}

	return false
}

func (c *Chain) GetChainAssets() []*Asset {
	return c.Assets
}
