package home

import (
	"context"

	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/home/entities"

	searchEntities "github.com/Jacobo0312/microservice-latency-experiment/internal/domain/search/entities"
	userEntities "github.com/Jacobo0312/microservice-latency-experiment/internal/domain/user/entities"
)

func (s *service) GetSectionsWithoutConcurrency(ctx context.Context, params entities.HomeParams) (*entities.Home, error) {

	searchData, err := s.SearchService.Get(ctx, searchEntities.Params{
		UserID: params.UserID,
	})

	if err != nil {
		return nil, err
	}

	userData, err := s.UserService.Get(ctx, userEntities.Params{
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
		UserSection: &entities.UserSection{
			Payload: userData.Payload,
			Name:    "get_user",
			Result:  "success",
		},
	}

	return &entities.Home{
		HomeData: entities.HomeData{
			Sections: &sections,
		},
	}, nil

}
