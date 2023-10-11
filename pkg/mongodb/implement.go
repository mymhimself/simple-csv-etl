package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// InsertOne inserts a single document into a MongoDB collection.
func (s *iMongoDB) InsertOne(ctx context.Context, database string, collection string, document interface{}) error {
	coll := s.client.Database(database).Collection(collection)
	_, err := coll.InsertOne(ctx, document)
	return err
}

// ─────────────────────────────────────────────────────────────────────────────
// FindOne finds a single document in a MongoDB collection based on a filter.
func (s *iMongoDB) FindOne(ctx context.Context, database string, collection string, filter interface{}) (any, error) {
	coll := s.client.Database(database).Collection(collection)
	var result bson.M
	err := coll.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// UpdateOne updates a single document in a MongoDB collection based on a filter and update data.
func (s *iMongoDB) UpdateOne(ctx context.Context, database string, collection string, filter any, update any) error {
	panic("unimplemented")

}

// ─────────────────────────────────────────────────────────────────────────────
// DeleteOne deletes a single document in a MongoDB collection based on a filter.
func (s *iMongoDB) DeleteOne(ctx context.Context, database string, collection string, filter interface{}) error {
	panic("unimplemented")
}

// ─────────────────────────────────────────────────────────────────────────────
// FindMany implements IMongoDB.
func (s *iMongoDB) FindMany(ctx context.Context, database string, collection string, filter any) ([]map[string]string, error) {
	coll := s.client.Database(database).Collection(collection)
	var results []map[string]string

	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var doc map[string]string

		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		results = append(results, doc)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
