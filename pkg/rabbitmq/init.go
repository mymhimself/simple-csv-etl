package rabbitmq

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

type rabbit struct {
	sync.Mutex
	wg sync.WaitGroup

	conn                 IRabbitConn //*amqp.Connection
	chann                IRabbitChan //*amqp.Channel
	consumerNames        []string
	resetConsumer        chan struct{}
	recreateConsumer     chan struct{}
	gracefulShutdown     chan struct{}
	NotifyConnCloseChan  chan *amqp.Error
	NotifyChannCloseChan chan *amqp.Error
	NotifyPublishChan    chan amqp.Confirmation
	NotifyBlockedChan    chan amqp.Blocking
}

func newp(ctx context.Context, url string, options ...Option) (*rabbit, error) {
	var rb = &rabbit{}
	rb.resetConsumer = make(chan struct{}, 100)
	err := rb.connectAndConfig(ctx, url, options...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rb.gracefulShutdown = make(chan struct{}, 10)
	go func() {
		var wait time.Duration = 0
		for {
			select {
			case <-rb.NotifyChannCloseChan:
				log.Println("Connection or Channel is lost, trying to re-establish...")
				//rb.Lock()
				time.Sleep(wait)
				err = rb.connectAndConfig(ctx, url, options...)
				//rb.Unlock()
				if err != nil {
					wait += 1 * time.Second
					log.Println(err)
				}
				if err == nil {
					log.Println("Connection Recoverd.")
					wait = 0
				}
			case <-rb.NotifyConnCloseChan:
				log.Println("Connection or Channel is lost, trying to re-establish...")
				//rb.Lock()
				time.Sleep(wait)
				err = rb.connectAndConfig(ctx, url, options...)
				//rb.Unlock()
				if err != nil {
					log.Println(err)
					wait += 1 * time.Second
				}
				if err == nil {
					log.Println("Connection Recoverd.")
					wait = 0
				}

			case <-ctx.Done():
				rb.GracefulShutdown()
				return
			}
		}
	}()

	return rb, nil
}

func (rb *rabbit) connectAndConfig(ctx context.Context, url string, options ...Option) error {
	rb.Lock()
	defer rb.Unlock()
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Println(err)
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
		return err
	}

	rb.conn = conn
	rb.chann = ch

	if len(options) > 0 {
		for _, f := range options {
			if err := f(rb.chann); err != nil {
				return err
			}
		}
	}

	rb.NotifyConnCloseChan = rb.conn.NotifyClose(make(chan *amqp.Error))
	rb.NotifyBlockedChan = rb.conn.NotifyBlocked(make(chan amqp.Blocking))
	rb.NotifyChannCloseChan = rb.chann.NotifyClose(make(chan *amqp.Error))
	rb.NotifyPublishChan = rb.chann.NotifyPublish(make(chan amqp.Confirmation, 1))
	if err = rb.chann.Confirm(false); err != nil {
		return err
	}
	if len(rb.consumerNames) > 0 {
		log.Println("a signal was sent on resetConsumer channel")
		rb.resetConsumer <- struct{}{}
	}

	return nil
}
