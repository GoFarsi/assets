package assets

import (
	"errors"
	"github.com/GoFarsi/assets/entity"
)

// getPaginatedChainList returns paginated list of chains with pageNumber and pageSize params
func getPaginatedChainList(chains []*entity.Chain, pageNumber, pageSize int) (result []*entity.Chain, err error) {
	totalChains := len(chains)
	pageNumber, err = validatePageNumber(pageNumber)
	if err != nil {
		return nil, err
	}

	start := (pageNumber - 1) * pageSize
	if start > totalChains {
		return nil, nil
	}

	end := start + pageSize
	if end > totalChains {
		end = totalChains
	}

	return chains[start:end], nil
}

// validatePageNumber will check the value of pageNumber to be greater than 0
func validatePageNumber(pageNumber int) (int, error) {
	if pageNumber == 0 {
		return 0, errors.New("pageNumber is invalid")
	}

	return pageNumber, nil
}
