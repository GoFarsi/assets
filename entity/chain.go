package entity

type ChainType string

const (
	MainChainType ChainType = "main"
	TestChainType           = "test"
)

type Chain struct {
	Id       string            `yaml:"id"`
	ParentId string            `yaml:"parent_id"`
	Name     string            `yaml:"name"`
	Symbol   string            `yaml:"symbol"`
	LogoPNG  string            `yaml:"logo_png"`
	ChainId  string            `yaml:"chain_id"`
	Assets   map[string]*Asset `yaml:"assets"`
	Type     ChainType         `yaml:"type"`
}

func NewChain(Id, name, symbol string, chainType ChainType, chainId string) *Chain {
	return &Chain{
		Id:      Id,
		Name:    name,
		Symbol:  symbol,
		ChainId: chainId,
		Type:    chainType,
	}
}

func (c *Chain) GetId() string {
	return c.Id
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

func (c *Chain) GetChainAssets() map[string]*Asset {
	return c.Assets
}
