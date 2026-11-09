package utils

// wire everything from the utilities -> server
import (
	"godiddy/internal/adapters/response"
	"godiddy/internal/core/domain"
)

func Pass(h *domain.Header, q *[]domain.Question) (*domain.Response, error){
	r, err := response.ConstructResponse(h, q)

	if err != nil {
		return nil, err
	}

	return r, nil
}
