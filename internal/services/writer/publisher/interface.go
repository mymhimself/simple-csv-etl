package publisher

import (
	"context"
)

type IPublisher interface {
	PublishCreate(ctx context.Context, object map[string]string) error
}
