package etl

import (
	"github.com/mymhimself/simple-csv-reader/internal/services/writer"
	"github.com/mymhimself/simple-csv-reader/pkg/config"
)

type iCollections struct {
	service writer.IWriter
}

func New(ops ...InitOption) (ICollections, error) {
	s := new(iCollections)
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

	return s, nil
}
