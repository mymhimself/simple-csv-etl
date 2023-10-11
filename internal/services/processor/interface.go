package processor

import "context"

type IProcessor interface {
	ProcessLines(ctx context.Context, lineChan chan string) error
}
