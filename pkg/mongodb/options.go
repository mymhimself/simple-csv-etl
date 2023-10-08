package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InitOption func(*iMongoDB) error

func InitOptionURI(v string) InitOption {
	return func(s *iMongoDB) error {
		// ctx will be used to set deadline for process, here
		// deadline will of 30 seconds.
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

		// mongo.Connect return mongo.Client method
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(v))
		if err != nil {
			return err
		}

		s.client = client
		return nil
	}
}
