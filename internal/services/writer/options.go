package writer

import "github.com/mymhimself/simple-csv-reader/internal/models/writer"

type InitOption func(*iWriter) error

func InitOptionModel(v writer.IWriter) InitOption {
	return func(s *iWriter) error {
		s.model = v
		return nil
	}
}
