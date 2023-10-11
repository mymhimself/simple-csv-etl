package publisher

import (
	"context"
)

type IPublisher interface {
	CreateNewRecord(ctx context.Context, params *CreateNewRecordParams) error
}

// ─────────────────────────────────────────────────────────────────────────────
type CreateNewRecordParams struct {
	Object     map[string]string
	Collection string
	Database   string
}
