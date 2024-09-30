package home

import (
	"context"

	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/home/entities"
)

type (
	Service interface {
		GetSectionsWithoutConcurrency(ctx context.Context, params entities.HomeParams) (*entities.Home, error)
		GetSectionsWithConcurrency(ctx context.Context, params entities.HomeParams) (*entities.Home, error)
	}
	service struct {
		*Container
	}
)

func New(c *Container) Service {
	return &service{c}
}
