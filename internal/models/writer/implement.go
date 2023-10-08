package writer

import (
	"context"
)

// ─────────────────────────────────────────────────────────────────────────────
// Create implements IWriter.
func (s *iWriter) Create(ctx context.Context, params *CreateParams) error {
	return s.mongodb.InsertOne(ctx, s.database, params.Collection, params.Object)
}

// ─────────────────────────────────────────────────────────────────────────────
// List implements IWriter.
func (s *iWriter) List(ctx context.Context, params *ListParams) ([]map[string]string, error) {
	result, err := s.mongodb.FindMany(ctx, s.database, params.Collection, nil)
	if err != nil {
		return nil, err
	}

	m, ok := result.([]map[string]string)
	if !ok {
		return nil, ErrInvalidInterfaceAssertion
	}

	return m, nil
}
