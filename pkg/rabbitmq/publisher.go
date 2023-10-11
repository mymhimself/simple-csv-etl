package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func NewPublisher(ctx context.Context, url string, options ...Option) (IRabbitPublisher, error) {
	rb, err := newp(ctx, url, options...)
	if err != nil {
		return rb, err
	}

	return rb, err
}

func (rb *rabbit) Publish(exchange string, key string, mandatory bool, immediate bool, msg amqp.Publishing) error {

	var err error
	for i := 1; i <= publishingRetryCount; i++ {
		err = rb.chann.Publish(
			exchange,
			key,
			mandatory,
			immediate,
			msg,
		)
		if err != nil {
			log.Println(err.Error()+". Retrying #", i)
			time.Sleep(time.Second * publishingTimeSecSleep * time.Duration(i))
			continue
		}
		if confirmed, ok := <-rb.NotifyPublishChan; !confirmed.Ack && ok {
			err = fmt.Errorf("publishing can not reach the RabbitMQ server")
			log.Println(err.Error()+". Retrying #", i)
			time.Sleep(time.Second * publishingTimeSecSleep * time.Duration(i))
			continue
		}
		return nil
	}
	//! todo send msg to a dead letter queue
	return err
}
