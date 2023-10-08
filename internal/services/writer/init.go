package writer

import (
	"github.com/mymhimself/simple-csv-reader/internal/models/writer"
	"github.com/mymhimself/simple-csv-reader/pkg/config"
	"github.com/mymhimself/simple-csv-reader/pkg/source"
)

type iWriter struct {
	model writer.IWriter

	consumer         source.Source
	consumersHandler map[string]consumerHandler
	config           struct {
		rmq struct {
			host         string
			exchangeName string
			queueName    string
			key          string
			consumerName string
		}
		numberOfConsumingThread int
	}
}

// ─────────────────────────────────────────────────────────────────────────────
func New(ops ...InitOption) (IWriter, error) {
	s := new(iWriter)

	for _, fn := range ops {
		err := fn(s)
		if err != nil {
			return nil, err
		}
	}

	// init consuming function
	s.consumersHandler = map[string]consumerHandler{}

	err := config.ValidateStruct(s)
	if err != nil {
		return nil, err
	}

	// init queue consumer
	err = s.setupConsumer()
	if err != nil {
		return nil, err
	}

	return s, nil
}
