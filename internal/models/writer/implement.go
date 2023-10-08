package writer

import (
	"context"
	"errors"
)

// ─────────────────────────────────────────────────────────────────────────────
// Create implements IWriter.
func (s *iWriter) Create(ctx context.Context, params *CreateParams) error {
	return s.mongodb.InsertOne(ctx, params.Collection, params.Object)
}

// ─────────────────────────────────────────────────────────────────────────────
// List implements IWriter.
func (s *iWriter) List(ctx context.Context, params *ListParams) ([]map[string]string, error) {
	result, err := s.mongodb.FindMany(ctx, "", nil)
	if err != nil {
		return nil, err
	}

	m, ok := result.([]map[string]string)
	if !ok {
		return nil, errors.New("invalid interface assertion")
	}

	return m, nil
}
