package publisher

import (
	"context"
	"encoding/json"
	"time"

	"github.com/streadway/amqp"
)

const contentJson = "application/json"

// ─────────────────────────────────────────────────────────────────────────────
// PublishCreate implements IPublisher
func (s *iPublisher) PublishCreate(ctx context.Context, object map[string]string) error {
	msg := &utilityGenericMessage{
		RUID:    time.Now().Format(time.RFC3339),
		Event:   "evt_create_new_record",
		Payload: object,
	}

	bs, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	err = s.publisher.Publish(s.config.exchange, s.config.key, false, false, amqp.Publishing{
		ContentType: contentJson,
		Timestamp:   time.Now(),
		Body:        bs,
	})
	if err != nil {
		return err
	}
	return nil
}

// ─────────────────────────────────────────────────────────────────────────────
