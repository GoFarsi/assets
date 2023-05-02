package assets

import "github.com/GoFarsi/assets/entity"

type Asset struct {
	chains []*entity.Chain
}

func New(assetsRepo string) *Asset {
	chains := parseAssetsByteToArray()
	asset := &Asset{chains: chains}

	switch assetsRepo {
	case GithubRepoUrl:
		asset.AdaptLogoPath(formatFilePathByGithubRepo)
	}

	return asset
}

func (a *Asset) GetTotalChainsSize() int {
	return len(a.chains)
}
