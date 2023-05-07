package assets

import (
	"errors"
	"github.com/GoFarsi/assets/entity"
)

// GetPaginatedChainList returns paginated list of chains with pageNumber and pageSize params
func GetPaginatedChainList(chains []*entity.Chain, pageNumber, pageSize int) (result []*entity.Chain, totalChains int, err error) {
	totalChains = len(chains)
	pageNumber, err = validatePageNumber(pageNumber)
	if err != nil {
		return nil, totalChains, err
	}

	start := (pageNumber - 1) * pageSize
	if start > totalChains {
		return nil, totalChains, nil
	}

	end := start + pageSize
	if end > totalChains {
		end = totalChains
	}

	return chains[start:end], totalChains, nil
}

// validatePageNumber will check the value of pageNumber to be greater than 0
func validatePageNumber(pageNumber int) (int, error) {
	if pageNumber == 0 {
		return 0, errors.New("pageNumber is invalid")
	}

	return pageNumber, nil
}
