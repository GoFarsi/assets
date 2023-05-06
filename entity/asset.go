package entity

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
