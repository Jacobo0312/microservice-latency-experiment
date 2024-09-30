package reader

import "github.com/Jacobo0312/microservice-latency-experiment/internal/rest"

type handler struct {
	*Container
}

func New(c *Container) rest.Handler {
	return &handler{
		Container: c,
	}
}
