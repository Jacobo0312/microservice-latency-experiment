package reader

import (
	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/home/entities"
)

type HomeParams struct {
	UserID string `json:"user_id"`
}

func (p *HomeParams) ToDomain() entities.HomeParams {
	return entities.HomeParams{
		UserID: p.UserID,
		Sections: []string{
			"search",
		},
	}
}
