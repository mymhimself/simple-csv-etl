package processor

import (
	"github.com/mymhimself/simple-csv-reader/internal/services/writer/publisher"
	"github.com/mymhimself/simple-csv-reader/pkg/config"
)

type iProcessor struct {
	writerPublisher publisher.IPublisher
	object          map[string]string

	config struct {
		delimiter string
	}
}

// ─────────────────────────────────────────────────────────────────────────────
func New(ops ...InitOption) (IProcessor, error) {
	s := new(iProcessor)

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
