package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// iMongoDB represents a MongoDB database connection.
type iMongoDB struct {
	client *mongo.Client
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
	return s, nil
}
