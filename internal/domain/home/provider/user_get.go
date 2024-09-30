package provider

import (
	"context"

	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/user"
	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/user/entities"
)

type getUser struct {
	srv  user.Service
	err  error
	data *string
	name string
}

func (s *getUser) Get(ctx context.Context, userID string) {
	data, err := s.srv.Get(ctx, entities.Params{
		UserID: userID,
	})

	s.data = &data.Payload
	s.err = err
}

func (s *getUser) Data(value interface{}) error {
	if s.err != nil {
		return s.err
	}
	if s.data != nil {
		v := value.(*string)
		*v = *s.data
	}
	return nil
}

func (c *getUser) Name() string {
	return c.name
}

func NewUserGet(srv user.Service) Providers {
	return &getUser{
		name: "get_user",
		srv:  srv,
	}
}
