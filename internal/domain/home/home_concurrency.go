package home

import (
	"context"
	"sync"

	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/home/entities"
	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/home/provider"
	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/home/section"
)

func (s *service) GetSectionsWithConcurrency(ctx context.Context, params entities.HomeParams) (*entities.Home, error) {
	providersCfg := map[string]section.Sections{
		"search": section.NewSearch(
			getSearchSectionParams(),
			provider.NewSearchGet(s.SearchService),
		),
		"user": section.NewUser(
			getUserSectionParams(),
			provider.NewUserGet(s.UserService),
		),
	}

	providers := make(map[string]provider.Providers)
	sections := make(map[string]section.Sections)

	for _, sectionValue := range params.Sections {
		sections[sectionValue] = providersCfg[sectionValue]
		for _, providerValue := range providersCfg[sectionValue].Providers() {
			providers[providerValue.Name()] = providerValue
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(providers))
	for _, providerValue := range providers {
		go func(providerValue provider.Providers) {
			defer wg.Done()
			providerValue.Get(ctx, params.UserID)
		}(providerValue)
	}
	wg.Wait()

	result := &entities.Home{
		HomeData: entities.HomeData{
			// UserID:   params.UserID,
			Sections: &entities.Sections{},
		},
	}

	var wgSections sync.WaitGroup
	wgSections.Add(len(sections))
	for _, sectionValue := range sections {
		go func(sectionValue section.Sections) {
			defer wgSections.Done()
			sectionValue.SetProviders(providers)
			sectionValue.Build(ctx, result)
		}(sectionValue)
	}
	wgSections.Wait()

	return result, nil
}

func getSearchSectionParams() entities.SearchSectionParams {
	return entities.SearchSectionParams{}
}

func getUserSectionParams() entities.UserSectionParams {
	return entities.UserSectionParams{}
}
