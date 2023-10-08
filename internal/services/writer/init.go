package writer

import (
	"github.com/mymhimself/simple-csv-reader/internal/models/writer"
)

type iWriter struct {
	model writer.IWriter
}

func New(ops ...InitOption) (IWriter, error) {
	s := new(iWriter)

	for _, fn := range ops {
		err := fn(s)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}
