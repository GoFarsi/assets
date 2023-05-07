package assets

import "github.com/GoFarsi/assets/entity"

type Asset struct {
	Chains []*entity.Chain
}

func New(assetsRepo string) *Asset {
	chains := parseAssetsByteToArray()
	asset := &Asset{Chains: chains}

	switch assetsRepo {
	case GithubRepoUrl:
		asset.AdaptLogoPath(formatFilePathByGithubRepo)
	}

	return asset
}

// GetTotalChainsSize return len of all chains in assets.yaml
func (a *Asset) GetTotalChainsSize() int {
	return len(a.Chains)
}

// GetTestChains return test chains (networks) in assets.yaml
func (a *Asset) GetTestChains() []*entity.Chain {
	return getChainsByType(a.Chains, entity.TestChainType)
}

// GetMainChains return main chains (networks) in assets.yaml
func (a *Asset) GetMainChains() []*entity.Chain {
	return getChainsByType(a.Chains, entity.MainChainType)
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
