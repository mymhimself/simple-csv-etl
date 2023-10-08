package reader

import "context"

type ICSVReader interface {
	ReadLines(ctx context.Context, lineChan chan string) error
}
