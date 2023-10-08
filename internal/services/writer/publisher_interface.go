package writer

import (
	"context"
)

type IPublisher interface {
	Create(ctx context.Context, object map[string]string) error
}
