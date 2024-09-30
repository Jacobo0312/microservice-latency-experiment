package section

import (
	"context"

	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/home/entities"
	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/home/provider"
)

const (
	_userSection string = "user"
)

type user struct {
	name          string
	providersData map[string]provider.Providers
	providers     []provider.Providers
}

func (s *user) Providers() []provider.Providers {
	return s.providers
}

func (s *user) SetProviders(providers map[string]provider.Providers) {
	s.providersData = providers
}

func NewUser(sectionParams entities.UserSectionParams, providers ...provider.Providers) Sections {
	return &user{
		name:      _userSection,
		providers: providers,
	}
}

func (s *user) Build(ctx context.Context, data *entities.Home) {
	s.initializeUserSection(data)

	err := s.providersData["get_user"].Data(&data.Sections.UserSection.Payload)

	if err != nil {
		data.Sections.UserSection.Result = "error"
		return
	}

	data.Sections.UserSection.Result = "success"
}

func (s *user) initializeUserSection(data *entities.Home) {
	data.Sections.UserSection = &entities.UserSection{
		Payload: "",
		Name:    _userSection,
	}
}
