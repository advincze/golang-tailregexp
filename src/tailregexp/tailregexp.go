package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"
)


func main() {

	
	file, err := os.Open("tail.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lineChan := make(chan string)
	go tailLine(file, lineChan)

	for {
		line := <-lineChan
		log.Printf("line: %s \n", line)
	}
}

func tailLine(f *os.File, lineChan chan string) {
	//go to the end of the log file
	f.Seek(0, 2)

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')

		if err !=nil {
			if err != io.EOF {
				log.Panic(err)
			}
		} else {
			lineChan <- line
		}

		time.Sleep(1e5)
	}
}


