package writer

import "github.com/mymhimself/simple-csv-reader/internal/services/writer/publisher"


package utility

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

)

type consumerHandler func(*genericMessage) error


// ─────────────────────────────────────────────────────────────────────────────
func (s *iWriter) OnMessage(bs []byte) error {
	var msg genericMessage
	err := json.Unmarshal(bs, &msg)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	fn, err := s.events(msg.Event)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	err = fn(&msg)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	return nil
}
// ─────────────────────────────────────────────────────────────────────────────
func (s *iPerson) events(event string) (consumerHandler, error) {
	var es = map[string]func(*genericMessage) error{
		"evt_create_new_record":       s.createNewRecord,
	}
	fn, ok := es[event]
	if !ok {
		return nil, ErrEventNotSupported
	}
	return fn, nil
}
// ─────────────────────────────────────────────────────────────────────────────
func (s *iWriter) loadDefaults() {
	s.config.numberOfConsumingThread = 1
	s.config.rmq.exchangeName = "ex_writer"
	s.config.rmq.queueName = "generic_writer"
	s.config.rmq.consumerName = "service_simple_csv_etl"
}
// ─────────────────────────────────────────────────────────────────────────────
func (s *iWriter) setupConsumer() error {
	var err error
	
	if s.consumer == nil {
		s.consumer, err = source.New(
			source.OptionWithHost(s.config.rmq.host),
			source.OptionWithExchangeName(s.config.rmq.exchangeName),
			source.OptionWithConsumerName(s.config.rmq.consumerName),
			source.OptionWithKey(s.config.rmq.key),
			source.OptionWithQueueName(s.config.rmq.queueName),
		)
		if err != nil {
			return err
		}
	}

	err = s.consumer.ForeachN(s.OnMessage, s.config.numberOfConsumingThread)
	if err != nil {
		return err
	}

	return nil
}
// ─────────────────────────────────────────────────────────────────────────────
func (s *iWriter) createNewRecord(msg *genericMessage) error {
	ctx := context.TODO()

	var params publisher.CreateNewRecordParams
	err := msg.unmarshalSubMessage(&params)
	if err != nil {
		return err
	}

	err := s.Create(ctx, &CreateParams{
		Collection: params.Collection,
		Object: params.Object,
	})
	if err != nil {
		sentry.CaptureException(err)
		return err
	}


	// delete all structure user have
	return nil
}
