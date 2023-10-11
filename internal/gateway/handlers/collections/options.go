package etl

import (
	"github.com/mymhimself/simple-csv-reader/internal/services/writer"
)

type InitOption func(s *iCollections) error

func InitOptionService(v writer.IWriter) InitOption {
	return func(s *iCollections) error {
		s.service = v
		return nil
	}
}
