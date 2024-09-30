package provider

import "context"

type Providers interface {
	Get(ctx context.Context, userID string)
	Data(value interface{}) error
	Name() string
}
