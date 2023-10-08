package writer

import (
	"context"
)

type IWriter interface {
	Create(ctx context.Context, object map[string]string) error
	List(ctx context.Context) ([]map[string]string, error)
}
