package rabbitmq

import (
	"context"

	"github.com/streadway/amqp"
)

type IRabbitConsumer interface {
	Consume(ctx context.Context, queue string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args amqp.Table, processor func(message *amqp.Delivery) error, handleError func(error)) error
	GracefulShutdown() error
}

type IRabbitPublisher interface {
	Publish(exchange string, key string, mandatory bool, immediate bool, msg amqp.Publishing) error
	CloseConnection() error
}

type IRabbitChan interface {
	ExchangeDeclare(name string, kind string, durable bool, autoDelete bool, internal bool, noWait bool, args amqp.Table) error
	QueueDeclare(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table) (amqp.Queue, error)
	QueueBind(name string, key string, exchange string, noWait bool, args amqp.Table) error
	Publish(exchange string, key string, mandatory bool, immediate bool, msg amqp.Publishing) error
	Consume(queue string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
	Qos(prefetchCount int, prefetchSize int, global bool) error
	NotifyPublish(confirm chan amqp.Confirmation) chan amqp.Confirmation
	NotifyClose(c chan *amqp.Error) chan *amqp.Error
	Confirm(noWait bool) error
	Cancel(consumer string, noWait bool) error
	Close() error
}

type IRabbitConn interface {
	Channel() (*amqp.Channel, error)
	NotifyClose(receiver chan *amqp.Error) chan *amqp.Error
	NotifyBlocked(receiver chan amqp.Blocking) chan amqp.Blocking
	Close() error
}
