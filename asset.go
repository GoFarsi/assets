package assets

import (
	"github.com/GoFarsi/assets/entity"
	"golang.org/x/exp/maps"
	"io"
	"net/http"
)

type AssetRepo struct {
	Chains map[string]*entity.Chain
}

type Pagination struct {
	PageNumber int
	PageSize   int
}

type Option struct {
	*Pagination
}

func New(repoAddress string) *AssetRepo {

	client := http.Client{}
	resp, err := client.Get(YamlAddress)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	assetsByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	chains := parseAssetsByteToArray(assetsByte)
	asset := &AssetRepo{Chains: chains}

	switch repoAddress {
	case GithubRepoUrl:
		asset.adaptLogoPath(formatFilePathByGithubRepo)
	}

	return asset
}

// GetTotalChainsSize return len of all chains in assets.yaml
func (a *AssetRepo) GetTotalChainsSize() int {
	return len(a.Chains)
}

// GetChains return all chains (networks) in assets.yaml
func (a *AssetRepo) GetChains(option *Option) ([]*entity.Chain, error) {
	return applyOptionsOnChains(a.Chains, option)
}

// GetTestChains return test chains (networks) in assets.yaml
func (a *AssetRepo) GetTestChains(option *Option) ([]*entity.Chain, error) {
	chains := getChainsByType(a.Chains, entity.TestChainType)
	return applyOptionsOnChains(chains, option)
}

// GetMainChains return main chains (networks) in assets.yaml
func (a *AssetRepo) GetMainChains(option *Option) ([]*entity.Chain, error) {
	chains := getChainsByType(a.Chains, entity.MainChainType)
	return applyOptionsOnChains(chains, option)
}

// GetChain return chain by its ID
func (a *AssetRepo) GetChain(Id string) *entity.Chain {
	return a.Chains[Id]
}

// GetChainBySymbol return chain by its symbol
func (a *AssetRepo) GetChainBySymbol(symbol string) *entity.Chain {
	for _, v := range a.Chains {
		if v.Symbol == symbol {
			return v
		}
	}

	return nil
}

// GetChainByName return chain by its name
func (a *AssetRepo) GetChainByName(name string) *entity.Chain {
	for _, v := range a.Chains {
		if v.Name == name {
			return v
		}
	}

	return nil
}

// GetAsset return asset by its ID
func (a *AssetRepo) GetAsset(Id string) *entity.Asset {
	for _, c := range a.Chains {
		if c.Assets[Id] != nil {
			return c.Assets[Id]
		}
	}

	return nil
}

// applyOptionsOnChains will check Options passed to requests and apply theme to result chains
func applyOptionsOnChains(chains map[string]*entity.Chain, option *Option) ([]*entity.Chain, error) {

	if option != nil && option.Pagination != nil {
		return getPaginatedChainList(chains, option.Pagination.PageNumber, option.Pagination.PageSize)
	}

	return maps.Values(chains), nil
}

// getChainsByType return list of chains by selecting type of chain (test, main,...)
func getChainsByType(chains map[string]*entity.Chain, chainType entity.ChainType) map[string]*entity.Chain {
	result := make(map[string]*entity.Chain)
	for k, v := range chains {
		if v.Type != chainType {
			continue
		}
		result[k] = v
	}

	return result
}
