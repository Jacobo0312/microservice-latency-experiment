package home

import (
	"context"

	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/home/entities"

	searchEntities "github.com/Jacobo0312/microservice-latency-experiment/internal/domain/search/entities"
)

func (s *service) GetSectionsWithoutConcurrency(ctx context.Context, params entities.HomeParams) (*entities.Home, error) {

	searchData, err := s.SearchService.Get(ctx, searchEntities.Params{
		UserID: params.UserID,
	})

	if err != nil {
		return nil, err
	}

	sections := entities.Sections{
		SearchSection: &entities.SearchSection{
			Payload: searchData.Payload,
			Name:    "get_search",
			Result:  "success",
		},
	}

	return &entities.Home{
		HomeData: entities.HomeData{
			Sections: &sections,
		},
	}, nil

}
