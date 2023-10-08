package mongodb

import (
	"context"
)

// MongoDB defines the interface for interacting with MongoDB.
type IMongoDB interface {
	InsertOne(ctx context.Context, document any) error
	FindOne(ctx context.Context, filter any) (any, error)
	UpdateOne(ctx context.Context, filter any, update any) error
	DeleteOne(ctx context.Context, filter any) error
}
