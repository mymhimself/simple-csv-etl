package businessdata

import "github.com/mymhimself/simple-csv-reader/internal/entities"

type IBusinessData interface {
	Create(bd *entities.BusinessData) error
}
