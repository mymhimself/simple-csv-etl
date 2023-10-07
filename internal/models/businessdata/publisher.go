package businessdata

import "github.com/mymhimself/simple-csv-reader/internal/entities"

type IPublisher interface {
	StoreBusinessData(bd *entities.BusinessData) error
}
