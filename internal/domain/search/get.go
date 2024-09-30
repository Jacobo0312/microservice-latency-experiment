package search

import (
	"context"
	"time"

	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/search/entities"
)

func (s *service) Get(ctx context.Context, params entities.Params) (*entities.Data, error) {

	//Simulate Latency
	time.Sleep(100 * time.Millisecond) // 50ms

	return &entities.Data{
		Payload: "Search Data",
	}, nil
}
