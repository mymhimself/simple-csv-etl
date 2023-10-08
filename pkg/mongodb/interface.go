package mongodb

import (
	"context"
)

// MongoDB defines the interface for interacting with MongoDB.
type IMongoDB interface {
	InsertOne(ctx context.Context, database string, collection string, document any) error
	FindOne(ctx context.Context, database string, collection string, filter any) (any, error)
	UpdateOne(ctx context.Context, database string, collection string, filter any, update any) error
	DeleteOne(ctx context.Context, database string, collection string, filter any) error
	FindMany(ctx context.Context, database string, collection string, filter any) (any, error)
}
