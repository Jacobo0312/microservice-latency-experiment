package user

import (
	"context"
	"time"

	"github.com/Jacobo0312/microservice-latency-experiment/internal/domain/user/entities"
)

func (s *service) Get(ctx context.Context, params entities.Params) (*entities.Data, error) {

	//Simulate Latency
	time.Sleep(150 * time.Millisecond) // 50ms

	return &entities.Data{
		Payload: "User Data",
	}, nil
}
