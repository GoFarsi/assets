package entity

import "strings"

type ContractType string
type LinkType string

const (
	PrimaryContractType ContractType = "primary"
	ProxyContractType                = "proxy"

	WebsiteLink    LinkType = "website"
	GithubLink              = "github"
	WhitePaperLink          = "whitepaper"
	TwitterLink             = "twitter"
	RedditLink              = "reddit"
	FacebookLink            = "facebook"
)

type Asset struct {
	Id          string                  `yaml:"id"`
	Name        string                  `yaml:"name"`
	Symbol      string                  `yaml:"symbol"`
	Contracts   map[ContractType]string `yaml:"contracts"`
	Decimals    int                     `yaml:"decimals"`
	MetaData    any                     `yaml:"metaData"`
	Description string                  `yaml:"description"`
	Issuer      string                  `yaml:"issuer"`
	Links       map[LinkType]string     `yaml:"links"`
	LogoPNG     string                  `yaml:"logo_png"`
	Types       []string                `yaml:"types"`
	Standards   []string                `yaml:"standards"`
}

func (a *Asset) GetId() string {
	return a.Id
}

func (a *Asset) GetName() string {
	return a.Name
}

func (a *Asset) GetSymbol() string {
	return a.Symbol
}

func (a *Asset) GetLogoPNG() string {
	return a.LogoPNG
}

func (a *Asset) GetPrimaryContractAddress() string {
	return a.Contracts[PrimaryContractType]
}

func (a *Asset) GetProxyContractAddress() string {
	return a.Contracts[ProxyContractType]
}

func (a *Asset) GetDecimals() int {
	return a.Decimals
}

func (a *Asset) GetDescription() string {
	return a.Description
}

func (a *Asset) HasErc20() bool {
	for _, s := range a.Standards {
		if strings.ToUpper(s) == "ERC20" {
			return true
		}
	}

	return false
}

func (a *Asset) HasErc721() bool {
	for _, s := range a.Standards {
		if strings.ToUpper(s) == "ERC721" {
			return true
		}
	}

	return false
}
