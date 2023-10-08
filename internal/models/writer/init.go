package writer

import (
	"github.com/mymhimself/simple-csv-reader/pkg/mongodb"
)

type iWriter struct {
	mongodb mongodb.IMongoDB
}

func New() (IWriter, error) {
	s := new(iWriter)
	return s, nil
}
