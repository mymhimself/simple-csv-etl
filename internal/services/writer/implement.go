package writer

import (
	"context"

	"github.com/mymhimself/simple-csv-reader/internal/models/writer"
)

// ─────────────────────────────────────────────────────────────────────────────
// Create implements IWriter.
func (s *iWriter) Create(ctx context.Context, params *CreateParams) error {
	return s.model.Create(ctx, &writer.CreateParams{
		Collection: params.Collection,
		Object:     params.Object,
	})
}

// ─────────────────────────────────────────────────────────────────────────────
// List implements IWriter.
func (s *iWriter) List(ctx context.Context, params *ListParams) ([]map[string]string, error) {
	return s.model.List(ctx, &writer.ListParams{Collection: params.Collection})
}
