package writer

import "github.com/mymhimself/simple-csv-reader/pkg/mongodb"

type InitOption func(*iWriter) error

func InitOptionMongoClient(v mongodb.IMongoDB) InitOption {
	return func(s *iWriter) error {
		s.mongodb = v
		return nil
	}
}
