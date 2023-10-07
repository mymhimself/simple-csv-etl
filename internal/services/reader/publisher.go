package reader

type IPublisher interface {
	processLine(line string) error
}
