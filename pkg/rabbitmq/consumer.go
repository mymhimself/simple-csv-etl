package rabbitmq

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mymhimself/logger"

	"github.com/streadway/amqp"
)

type consumer struct {
	sync.Mutex
	wg     sync.WaitGroup
	config struct {
		ops []Option
		url string

		queueName    string
		consumerName string
		autoAck      bool
		exclusive    bool
		noLocal      bool
		noWait       bool
		args         amqp.Table
	}
	proccesor    func(message *amqp.Delivery) error
	errorHandler func(error)

	conn  *amqp.Connection //*amqp.Connection
	chann *amqp.Channel    //*amqp.Channel
}

func NewConsumer(ctx context.Context, url string, options ...Option) (IRabbitConsumer, error) {
	rb := new(consumer)

	rb.config.ops = options
	rb.config.url = url

	rb.Lock()
	defer rb.Unlock()
	err := rb.connect()
	if err != nil {
		return nil, err
	}
	rb.config.ops = options

	return rb, nil
}

func (rb *consumer) connect() error {
	conn, err := amqp.Dial(rb.config.url)
	if err != nil {
		return err
	}
	rb.conn = conn
	notifyConnBlock := conn.NotifyBlocked(make(chan amqp.Blocking, 1))
	notifyConnClose := conn.NotifyClose(make(chan *amqp.Error, 1))
	//wathout the connection
	go func() {
		select {
		case <-notifyConnBlock:
			var err error
			for i := 0; i < 10; i++ {
				time.Sleep(time.Duration(i*3) * time.Second)
				err = rb.connect()
				if err != nil {
					continue
				}
				if err == nil {
					return
				}
			}
			if err != nil {
				logger.Error("Couldn't recover the connection")
				panic(err)
			}
		case qerr := <-notifyConnClose:
			logger.Error(qerr)
			var err error
			for i := 0; i < 10; i++ {
				time.Sleep(time.Duration(i*3) * time.Second)
				err = rb.connect()
				if err != nil {
					continue
				}
				if err == nil {
					return
				}
			}
			if err != nil {
				logger.Error("Couldn't recover the connection")
				panic(err)
			}
		}
	}()
	return nil
}

func (rb *consumer) Consume(ctx context.Context, queue string, consumerName string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args amqp.Table, processor func(message *amqp.Delivery) error, handleError func(error)) error {
	if consumerName == "" {
		return fmt.Errorf("a unique consumer name (scoped for all consumers on this channel) should be provided for the consumer argument")
	}
	logger.Debugf("Consuming started")
	rb.proccesor = processor
	rb.errorHandler = handleError

	rb.config.queueName = queue
	rb.config.consumerName = consumerName
	rb.config.autoAck = autoAck
	rb.config.exclusive = exclusive
	rb.config.noLocal = noLocal
	rb.config.noWait = noWait
	rb.config.args = args

	rb.wg.Add(1)
	go func() {
		// this is going to lock the thread.
		err := rb.consume()
		if err != nil {
			panic(err)
		}
	}()
	return nil
}

func (rb *consumer) consume() error {
	var err error
	if rb.conn.IsClosed() {
		// wait untill connection is open
		logger.Debugf("connection is not open, waiting for connection to be recoverd")
		time.Sleep(1 * time.Second)
		return rb.consume()
	}

	rb.chann, err = rb.conn.Channel()
	if err != nil {
		logger.Error(err)
		logger.Debugf("channel creation failed, waiting for connection to be recoverd")
		time.Sleep(1 * time.Second)
		return rb.consume()
	}

	for _, fn := range rb.config.ops {
		err := fn(rb.chann)
		if err != nil {
			panic(err)
		}
	}
	// rb.chann.Confirm(true)

	C, err := rb.chann.Consume(
		rb.config.queueName,
		rb.config.consumerName,
		rb.config.autoAck,
		rb.config.exclusive,
		rb.config.noLocal,
		rb.config.noWait,
		rb.config.args,
	)
	if err != nil {
		logger.Error(err)
		logger.Debugf("channel creation failed, waiting for connection to be recoverd")
		time.Sleep(1 * time.Second)
		return rb.consume()
	}

	for {
		d, ok := <-C
		if !ok {
			// channel is closed,
			return rb.consume()
		}
		err := rb.proccesor(&d)
		if err != nil && rb.errorHandler != nil {
			rb.errorHandler(err)
		}
	}

	return nil
}

// GracefulShutdown can be called by a consumer to gracfully shutdown after getting all deliveris processed
func (rb *consumer) GracefulShutdown() error {

	// err := rb.chann.Cancel(rb.config.consumerName, false)
	// if err != nil {
	// 	return err
	// }

	// rb.wg.Wait()
	// if rb.conn != nil {
	// 	if err := rb.conn.Close(); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}
