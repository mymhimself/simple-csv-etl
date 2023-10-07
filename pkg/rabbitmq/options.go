package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Option func(IRabbitChan) error

func OptionExchangeDeclare(name string, kind string, durable bool, autoDelete bool, internal bool, noWait bool, args amqp.Table) Option {
	return func(rb IRabbitChan) error {
		err := rb.ExchangeDeclare(
			name,
			kind,
			durable,
			autoDelete,
			internal,
			noWait,
			args,
		)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}
}

func OptionQueueDeclare(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table) Option {
	return func(rb IRabbitChan) error {
		_, err := rb.QueueDeclare(
			name,
			durable,
			autoDelete,
			exclusive,
			noWait,
			args,
		)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}
}

func OptionQueueBind(name string, key string, exchange string, noWait bool, args amqp.Table) Option {
	return func(rb IRabbitChan) error {
		err := rb.QueueBind(
			name,
			key,
			exchange,
			noWait,
			args,
		)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}
}

func OptionQos(prefetchCount int, prefetchSize int, global bool) Option {
	return func(rb IRabbitChan) error {
		err := rb.Qos(prefetchCount, prefetchSize, global)
		if err != nil {
			return err
		}
		return nil
	}
}
