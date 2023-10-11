package source

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/mymhimself/logger"
	rmq "github.com/mymhimself/simple-csv-reader/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

type rmqSource struct {
	c      rmq.IRabbitConsumer
	C      chan amqp.Delivery
	config struct {
		postfix      string
		host         string
		name         string
		queueName    string
		exchangeName string
		key          string
		bufferSize   int
		xExpires     int
		args         map[string]interface{}
	}
}

// ─────────────────────────────────────────────────────────────────────────────
func New(ops ...Option) (Source, error) {
	s := new(rmqSource)
	var err error

	s.config.postfix = "-" + uuid.NewString()[:5]
	// s.config.args = make(map[string]interface{})

	// s.config.args["x-expires"] = s.config.xExpires

	for _, fn := range ops {
		err = fn(s)
		if err != nil {
			return nil, err
		}
	}

	s.C = make(chan amqp.Delivery)
	rops := []rmq.Option{
		rmq.OptionQos(0, 0, false),
	}
	if s.config.exchangeName != "" {
		rops = append(
			rops,
			rmq.OptionQueueDeclare(s.GetQueueName(), true, false, false, false, s.config.args),
			rmq.OptionQueueBind(s.GetQueueName(), s.config.key, s.config.exchangeName, false, nil),
		)
	}
	s.c, err = rmq.NewConsumer(context.Background(), s.config.host,
		rops...,
	)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	// logger.Debugf("Source Created:%+v", s)
	return s, nil
}

// ─────────────────────────────────────────────────────────────────────────────
func (s *rmqSource) ForeachN(fn func([]byte) error, n int) error {
	var err error
	if s.c == nil {
		return ErrChannelNameInvalid
	}
	err = s.c.Consume(
		context.Background(),
		s.GetQueueName(),
		s.GetConsumerName(),
		false, false, false, false, nil,
		func(message *amqp.Delivery) error {
			s.C <- *message
			return nil
		}, func(err error) {
			sentry.CaptureException(err)
			logger.Error(err)
			sentry.CaptureException(err)
		},
	)
	if err != nil {
		return err
	}

	for i := 0; i < n; i++ {
		go func() {
		InnerLoop:
			for msg := range s.C {
				err = fn(msg.Body)
				if err != nil {
					// send ACK
					logger.Error(err)
					sentry.CaptureException(err)
					continue InnerLoop
				}
				msg.Ack(true)
			}
		}()
	}
	return nil
}

// ─────────────────────────────────────────────────────────────────────────────
func (s *rmqSource) GetQueueName() string {
	return s.config.queueName
}

func (s *rmqSource) GetConsumerName() string {
	return s.config.name + s.config.postfix
}
