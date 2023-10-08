package processor

import (
	"context"
	"strings"

	"github.com/mymhimself/logger"
)

// ─────────────────────────────────────────────────────────────────────────────
func (s *iProcessor) ProcessLines(ctx context.Context, lineChan chan string) error {
	for line := range lineChan {
		rowMap := s.extractObjectFromLine(line)

		err := s.writerPublisher.PublishCreate(ctx, rowMap)
		if err != nil {
			logger.Error(err)
		}
	}
	return nil
}

// ─────────────────────────────────────────────────────────────────────────────
func (s *iProcessor) extractObjectFromLine(line string) map[string]string {
	record := make(map[string]string)
	var i int8
	array := strings.Split(line, s.config.delimiter)
	for key := range s.object {
		record[key] = array[i]
		i++
	}

	return record
}
