package writer

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// ─────────────────────────────────────────────────────────────────────────────
// Create implements IWriter.
func (s *iWriter) Create(ctx context.Context, params *CreateParams) error {
	return s.mongodb.InsertOne(ctx, s.database, params.Collection, params.Object)
}

// ─────────────────────────────────────────────────────────────────────────────
// List implements IWriter.
func (s *iWriter) List(ctx context.Context, params *ListParams) ([]map[string]string, error) {
	return s.mongodb.FindMany(ctx, s.database, params.Collection, bson.D{})
}
