package assets

import (
	"github.com/GoFarsi/assets/entity"
)

type Asset struct {
	Chains []*entity.Chain
}

type Pagination struct {
	PageNumber int
	PageSize   int
}

type Option struct {
	*Pagination
}

func New(assetsRepo string) *Asset {
	chains := parseAssetsByteToArray()
	asset := &Asset{Chains: chains}

	switch assetsRepo {
	case GithubRepoUrl:
		asset.adaptLogoPath(formatFilePathByGithubRepo)
	}

	return asset
}

// GetTotalChainsSize return len of all chains in assets.yaml
func (a *Asset) GetTotalChainsSize() int {
	return len(a.Chains)
}

// GetTestChains return test chains (networks) in assets.yaml
func (a *Asset) GetTestChains(option *Option) ([]*entity.Chain, error) {
	chains := getChainsByType(a.Chains, entity.TestChainType)
	return applyOptionsOnChains(chains, option)
}

// GetMainChains return main chains (networks) in assets.yaml
func (a *Asset) GetMainChains(option *Option) ([]*entity.Chain, error) {
	chains := getChainsByType(a.Chains, entity.MainChainType)
	return applyOptionsOnChains(chains, option)
}

// applyOptionsOnChains will check Options passed to requests and apply theme to result chains
func applyOptionsOnChains(chains []*entity.Chain, option *Option) ([]*entity.Chain, error) {

	if option.Pagination != nil {
		return getPaginatedChainList(chains, option.Pagination.PageNumber, option.Pagination.PageSize)
	}

	return chains, nil
}

// getChainsByType return list of chains by selecting type of chain (test, main,...)
func getChainsByType(chains []*entity.Chain, chainType entity.ChainType) (result []*entity.Chain) {
	for _, c := range chains {
		if c.Type != chainType {
			continue
		}
		result = append(result, c)
	}

	return result
}
