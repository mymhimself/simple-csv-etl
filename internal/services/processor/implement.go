package processor

import (
	"context"
	"strings"

	"github.com/mymhimself/logger"
	"github.com/mymhimself/simple-csv-reader/internal/services/writer"
)

type iProcessor struct {
	writerPublisher writer.IPublisher
	object          map[string]string

	config struct {
		delimiter string
	}
}

// ─────────────────────────────────────────────────────────────────────────────
func (s *iProcessor) ProcessLines(ctx context.Context, lineChan chan string) error {
	for line := range lineChan {
		rowMap := s.extractObjectFromLine(line)

		err := s.writerPublisher.Create(ctx, rowMap)
		if err != nil {
			logger.Error(err)
		}
	}
	return nil
}

// ─────────────────────────────────────────────────────────────────────────────
func (s *iProcessor) extractObjectFromLine(line string) map[string]string {
	newObject := make(map[string]string)
	var i int8
	array := strings.Split(line, s.config.delimiter)
	for key := range s.object {
		newObject[key] = array[i]
		i++
	}

	return newObject
}
