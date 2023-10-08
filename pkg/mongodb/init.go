package mongodb

import (
	"context"

	"github.com/mymhimself/logger"
	"github.com/mymhimself/simple-csv-reader/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// iMongoDB represents a MongoDB database connection.
type iMongoDB struct {
	client *mongo.Client `validate:"shallow"`
}

// NewMongoDatabase creates a new instance of MongoDatabase.
func New(ops ...InitOption) (IMongoDB, error) {
	s := new(iMongoDB)
	for _, fn := range ops {
		err := fn(s)
		if err != nil {
			return nil, err
		}
	}

	err := config.ValidateStruct(s)
	if err != nil {
		return nil, err
	}

	{
		err := s.ping()
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

// ─────────────────────────────────────────────────────────────────────────────
func (s *iMongoDB) ping() error {

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := s.client.Ping(context.Background(), readpref.Primary()); err != nil {
		return err
	}
	logger.Info("mongodb connected successfully")
	return nil
}
