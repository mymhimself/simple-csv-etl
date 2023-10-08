package publisher

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/mymhimself/simple-csv-reader/pkg/rabbitmq"
)

const (
	defaultExchangeType = "direct"
	defaultExchangeName = "ex_writer"
)

type iPublisher struct {
	publisher rabbitmq.IRabbitPublisher
	config    struct {
		host         string
		key          string
		exchangeType string
		exchange     string
	}
}

// ─────────────────────────────────────────────────────────────────────────────
func New(ops ...PublisherOption) (IPublisher, error) {
	s := new(iPublisher)
	for _, fn := range ops {
		err := fn(s)
		if err != nil {
			return nil, err
		}
	}
	err := s.init()
	if err != nil {
		return nil, err
	}
	return s, nil
}

// ─────────────────────────────────────────────────────────────────────────────
func (s *iPublisher) loadDefaults() {
	s.config.exchangeType = defaultExchangeType
	s.config.exchange = defaultExchangeName
}

// ─────────────────────────────────────────────────────────────────────────────
func (s *iPublisher) init() error {
	var err error
	s.publisher, err = rabbitmq.NewPublisher(
		context.Background(), s.config.host,
		rabbitmq.OptionExchangeDeclare(s.config.exchange, s.config.exchangeType, true, false, false, false, nil),
		rabbitmq.OptionQos(0, 0, false),
	)

	// s.scheduler, err := s

	if err != nil {
		return err
	}

	return nil
}

// ─────────────────────────────────────────────────────────────────────────────

type publisherGenericMessage struct {
	RUID    string                 `json:"rid,omitempty"`
	Event   string                 `json:"event,omitempty"`
	Payload interface{}            `json:"payload,omitempty"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
}

// ─────────────────────────────────────────────────────────────────────────────
func (m *publisherGenericMessage) unmarshalSubMessage(v interface{}) error {
	if m.Payload == nil {
		return errors.New("nil payload")
	}
	bs, err := json.Marshal(m.Payload)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bs, v)
	if err != nil {
		return err
	}
	return nil
}
