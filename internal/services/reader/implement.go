package reader

import (
	"bufio"
	"context"
	"errors"
	"os"
	"strings"

	"github.com/mymhimself/simple-csv-reader/pkg/config"
)

type iCSVReader struct {
	object    map[string]string
	delimiter string
	fileName  string
}

// ─────────────────────────────────────────────────────────────────────────────
func New(ops ...InitOption) (ICSVReader, error) {
	s := new(iCSVReader)

	for _, fn := range ops {
		err := fn(s)
		if err != nil {
			return nil, err
		}
	}

	err := config.ValidateStruct(s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// ─────────────────────────────────────────────────────────────────────────────
func (s *iCSVReader) readMetadata() error {
	s.object = make(map[string]string)
	file, _ := os.Open(s.fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return errors.New("header not found, file is empty")
	}

	fieldsList := strings.Split(scanner.Text(), s.delimiter)
	for _, field := range fieldsList {
		s.object[field] = ""
	}

	return nil
}

// ─────────────────────────────────────────────────────────────────────────────
// StartReading implements ICSVReader.
// this function read each line and push it into the out put channel
func (s *iCSVReader) ReadLines(ctx context.Context, lineChan chan string) error {
	err := s.readMetadata()
	if err != nil {
		return err
	}

	file, _ := os.Open(s.fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// fake read for ignoring the header section
	scanner.Text()

	// reading line by line
	for scanner.Scan() {
		lineChan <- scanner.Text()
	}

	return nil
}
