package assets

import "github.com/GoFarsi/assets/entity"

type Asset struct {
	chains []*entity.Chain
}

func New() *Asset {
	chains := parseAssetsByteToArray()
	asset := &Asset{chains: chains}
	return asset
}

func (a *Asset) GetTotalChainsSize() int {
	return len(a.chains)
}
