package section

import (
	"context"

	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/home/entities"
	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/home/provider"
)

type Sections interface {
	Build(ctx context.Context, data *entities.Home)
	Providers() []provider.Providers
	SetProviders(providers map[string]provider.Providers)
}
