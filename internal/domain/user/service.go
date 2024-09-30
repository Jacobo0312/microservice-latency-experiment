package user

import (
	"context"

	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/user/entities"
)

type (
	Service interface {
		Get(ctx context.Context, params entities.Params) (*entities.Data, error)
	}
	service struct {
		*Container
	}
)

func New(c *Container) Service {
	return &service{c}
}
