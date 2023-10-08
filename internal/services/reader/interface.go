package reader

import "context"

type ICSVReader interface {
	ReadLines(ctx context.Context, lineChan chan string) error
	ReadMetaData(ctx context.Context) (map[string]string, error)
}
