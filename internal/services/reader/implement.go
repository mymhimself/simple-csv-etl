package reader

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/mymhimself/logger"
	"github.com/mymhimself/simple-csv-reader/internal/models/businessdata"
	"github.com/mymhimself/simple-csv-reader/pkg/config"
)

type iCSVReader struct {
	object    map[string]string
	publisher IPublisher

	config struct {
		delimiter string
		fileName  string
		model     businessdata.IBusinessData
	}
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
	file, _ := os.Open(s.config.fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return errors.New("header not found, file is empty")
	}

	fieldsList := strings.Split(scanner.Text(), s.config.delimiter)
	for _, field := range fieldsList {
		s.object[field] = ""
	}

	return nil
}

// ─────────────────────────────────────────────────────────────────────────────
// StartReading implements ICSVReader.
// this function read each line and push it into the out put channel
func (s *iCSVReader) ReadLines(lineChan chan string) error {
	err := s.readMetadata()
	if err != nil {
		return err
	}

	file, _ := os.Open(s.config.fileName)
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

// ─────────────────────────────────────────────────────────────────────────────
func (s *iCSVReader) ProcessLines(lineChan chan string) error {
	for line := range lineChan {
		err := s.publisher.processLine(line)
		if err != nil {
			logger.Error(err)
		}
	}

	return nil
}
