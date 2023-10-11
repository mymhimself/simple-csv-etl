package processor

import "github.com/mymhimself/simple-csv-reader/internal/services/writer/publisher"

type InitOption func(*iProcessor) error

func InitOptionPublisher(v publisher.IPublisher) InitOption {
	return func(s *iProcessor) error {
		s.writerPublisher = v
		return nil
	}
}

func InitOptionDelimiter(v string) InitOption {
	return func(s *iProcessor) error {
		s.config.delimiter = v
		return nil
	}
}

func InitOptionObject(v map[string]string) InitOption {
	return func(s *iProcessor) error {
		s.object = v
		return nil
	}
}

func InitOptionCollectionName(v string) InitOption {
	return func(s *iProcessor) error {
		s.collection = v
		return nil
	}
}
