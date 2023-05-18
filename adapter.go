package assets

import (
	"fmt"
	"github.com/GoFarsi/assets/entity"
	"sync"
)

type logoPathPattern func(path string) string

// adaptLogoPath will format all logo paths of chains and assets by using specific pattern
func (a *AssetRepo) adaptLogoPath(patternFunc logoPathPattern) {
	var wg sync.WaitGroup

	adaptedChains := make(chan *entity.Chain)
	for _, chain := range a.Chains {
		wg.Add(1)
		go func(chain *entity.Chain) {
			defer wg.Done()
			chain.LogoPNG = patternFunc(chain.LogoPNG)
			for _, asset := range chain.Assets {
				asset.LogoPNG = patternFunc(asset.LogoPNG)
			}
			adaptedChains <- chain
		}(chain)
	}

	go func() {
		wg.Wait()
		close(adaptedChains)
	}()

	for range adaptedChains {
	}
}

// formatPathToGithubMainRepo return a formatted path of file based on GitHub main branch
func formatFilePathByGithubRepo(path string) string {
	return fmt.Sprintf("%s/%s", GithubRepoUrl, path)
}
