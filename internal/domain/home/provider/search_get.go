package provider

import (
	"context"

	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/search"
	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/search/entities"
)

type getSearch struct {
	srv  search.Service
	err  error
	data *entities.Data
	name string
}

func (s *getSearch) Get(ctx context.Context, userID string) {
	data, err := s.srv.Get(ctx, entities.Params{
		UserID: userID,
	})

	s.data = data
	s.err = err
}

func (s *getSearch) Data(value interface{}) error {
	if s.err != nil {
		return s.err
	}
	if s.data != nil {
		v := value.(*entities.Data)
		*v = *s.data
	}
	return nil
}

func (c *getSearch) Name() string {
	return c.name
}

func NewSearchGet(srv search.Service) Providers {
	return &getSearch{
		name: "search",
		srv:  srv,
	}
}
