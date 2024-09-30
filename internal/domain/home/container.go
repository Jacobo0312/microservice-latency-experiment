package home

import (
	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/search"
	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/user"
)

type Container struct {
	SearchService            search.Service
	UserService              user.Service
	// RecommendedMoviesService recommendedMovies.Service
	// WatchedMoviesService     watchedMovies.Service
}
