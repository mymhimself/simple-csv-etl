package reader

type InitOption func(*iCSVReader) error

func InitOptionDelimiter(v string) InitOption {
	return func(s *iCSVReader) error {
		s.delimiter = v
		return nil
	}
}

func InitOptionFileName(v string) InitOption {
	return func(s *iCSVReader) error {
		s.fileName = v
		return nil
	}
}
