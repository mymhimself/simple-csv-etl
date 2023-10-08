package writer

import (
	"context"
)

type IWriter interface {
	Create(context.Context, map[string]string) error
	List(context.Context) []map[string]string
}
