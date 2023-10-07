package reader

type ICSVReader interface {
	ReadLines(lineChan chan string) error
	ProcessLines(lineChan chan string) error
}
