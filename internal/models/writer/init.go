package writer

import (
	"github.com/mymhimself/simple-csv-reader/pkg/config"
	"github.com/mymhimself/simple-csv-reader/pkg/mongodb"
)

type iWriter struct {
	mongodb  mongodb.IMongoDB
	database string
}

func New(ops ...InitOption) (IWriter, error) {
	s := new(iWriter)

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
