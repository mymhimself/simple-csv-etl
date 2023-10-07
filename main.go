package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	now := time.Now().UnixMilli()
	// readLineByLine()
	readAll()
	fmt.Println("\nElapsed Time:", time.Now().UnixMilli()-now)
}

func readLineByLine() {
	file, _ := os.Open("business-financial-data-mar-2022-quarter-csv.csv")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processLine(scanner.Text())
	}
}

func processLine(line string) {
	// fmt.Println(len([]byte(line)))
	fmt.Println(line)
}

func readAll() {
	content, err := os.ReadFile("business-financial-data-mar-2022-quarter-csv.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
