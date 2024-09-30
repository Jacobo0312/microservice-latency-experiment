package home

import (
	"context"
	"fmt"

	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/home/entities"

	searchEntities "github.com/Jacobo0312/microservice-latency-experiment/internal/domain/search/entities"
)

func (s *service) GetSections(ctx context.Context, params entities.HomeParams) (*entities.Home, error) {

	fmt.Println("Getting sections")

	searchData, err := s.SearchService.Get(ctx, searchEntities.Params{
		UserID: params.UserID,
	})

	if err != nil {
		return nil, err
	}

	sections := entities.Sections{
		SearchSection: &entities.SearchSection{
			Payload: searchData.Payload,
		},
	}

	return &entities.Home{
		Sections: sections,
	}, nil

}
