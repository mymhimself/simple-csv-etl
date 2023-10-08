package writer

import (
	"context"
)

type IWriter interface {
	Create(ctx context.Context, params *CreateParams) error
	List(ctx context.Context, params *ListParams) ([]map[string]string, error)
}

// ─────────────────────────────────────────────────────────────────────────────

type CreateParams struct {
	Collection string
	Object     map[string]string
}

// ─────────────────────────────────────────────────────────────────────────────
type ListParams struct {
	Collection string
}
