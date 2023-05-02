package assets

import (
	_ "embed"
	"github.com/GoFarsi/assets/entity"
	"gopkg.in/yaml.v3"
)

//go:embed resources/assets.yaml
var assetsByte []byte

// parseAssetsByteToArray parse the byte contents of yaml file to array of entity.Chain
func parseAssetsByteToArray() (chains []*entity.Chain) {
	_ = yaml.Unmarshal(assetsByte, &chains)
	return chains
}
