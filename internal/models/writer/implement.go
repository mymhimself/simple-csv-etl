package writer

import (
	"github.com/mymhimself/simple-csv-reader/pkg/mongodb"
)

type iWriter struct {
	mongodb mongodb.IMongoDB
}
