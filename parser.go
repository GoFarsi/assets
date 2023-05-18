package assets

import (
	_ "embed"
	"github.com/GoFarsi/assets/entity"
	"gopkg.in/yaml.v3"
)

// parseAssetsByteToArray parse the byte contents of yaml file to array of entity.Chain
func parseAssetsByteToArray(assetsByte []byte) (chains map[string]*entity.Chain) {
	_ = yaml.Unmarshal(assetsByte, &chains)
	return chains
}
