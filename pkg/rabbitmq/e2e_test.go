package rabbitmq

import (
	"encoding/base64"
	"log"
	"strconv"
	"testing"
	"time"

	"context"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/require"
)

const numberOfTestMessage = 100


func TestMain(m *testing.M) {
	// todo add the docker rmq

}

func testConsume(t *testing.T) {
	n := 0
	m := 0
	t.Parallel()
	C, err := NewConsumer(context.Background(), "amqp://guest:guest@localhost:5672/",
		OptionQos(0, 0, false),

		OptionQueueDeclare("test", false, false, false, false, nil),
	)

	require.NoError(t, err)
	err = C.Consume(context.Background(), "test", "test-consumer", false, false, false, false, nil, func(message *amqp.Delivery) error {
		n++
		if len(message.Body) < 1 {
			m++
		}
		log.Println("message resived")

		message.Ack(true)
		// require.NotEmpty(t, message.Body)
		return nil
	}, func(err error) {
		require.NoError(t, err)
	})
	require.NoError(t, err)

	time.Sleep(120 * time.Second)

	log.Printf("Consumed %d of %d\n", m, n)
}

func testPublish(t *testing.T) {
	t.Parallel()
	P, err := NewPublisher(context.Background(), "amqp://guest:guest@localhost:5672/", OptionQueueDeclare("test", false, false, false, false, nil), OptionQos(0, 0, false))
	require.NoError(t, err)
	defer func() {
		P.CloseConnection()
	}()
	for i := 0; i < numberOfTestMessage; i++ {
		err := P.Publish("", "test", false, false, amqp.Publishing{
			Type: "test",
			// DeliveryMode: ,
			Body: []byte(base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(i)))),
		})
		log.Println("message Published")
		time.Sleep(2 * time.Second)
		require.NoError(t, err)
	}
}
