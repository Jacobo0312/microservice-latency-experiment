package container

import (
	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/home"
	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/search"
	"github.com/Jacobo0312/microservice-latency-experiment/internal/rest"
	"github.com/Jacobo0312/microservice-latency-experiment/internal/rest/reader"
)

func Reader() rest.Handler {
	global := Container{
		Reader: &reader.Container{},
		Home:   &home.Container{},
	}

	setupReader(&global)

	return reader.New(global.Reader)
}

func setupReader(container *Container) {
	searchService := search.New(&search.Container{})
	container.Home.SearchService = searchService
	homeService := home.New(container.Home)
	container.Reader.HomeService = homeService

}
