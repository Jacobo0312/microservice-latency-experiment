package section

import (
	"context"

	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/home/entities"
	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/home/provider"
)

const (
	_searchSection string = "search"
)

type search struct {
	name          string
	providersData map[string]provider.Providers
	providers     []provider.Providers
}

func (s *search) Providers() []provider.Providers {
	return s.providers
}

func (s *search) SetProviders(providers map[string]provider.Providers) {
	s.providersData = providers
}

func NewSearch(sectionParams entities.SearchSectionParams, providers ...provider.Providers) Sections {
	return &search{
		name:      _searchSection,
		providers: providers,
	}
}

func (s *search) Build(ctx context.Context, data *entities.Home) {
	s.initializeSearchSection(data)

	err := s.providersData["get_search"].Data(&data.Sections.SearchSection.Payload)

	if err != nil {
		data.Sections.SearchSection.Result = "error"
		return
	}

	data.Sections.SearchSection.Result = "success"
}

func (s *search) initializeSearchSection(data *entities.Home) {
	data.Sections.SearchSection = &entities.SearchSection{
		Payload: "",
		Name:    _searchSection,
	}
}
