package container

import (
	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/home"
	"github.com/Jacobo0312/microservice-latency-experiment/internal/rest/reader"
)

type Container struct {
	Reader *reader.Container
	Home   *home.Container
}
