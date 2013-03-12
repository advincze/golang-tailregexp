package main

import (
	"bufio"
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"regexp"
	"time"
)

func main() {

	filePath := flag.String("file", "tail.txt", "file to watch")
	regex := flag.String("regex", `(.*)`, "regex to match")
	flag.Parse()
	println("logfile to watch: ", *filePath)
	println("regex to match: ", *regex)

	r, err := regexp.Compile(*regex)
	if err != nil {
		panic(err)
	}

	file, err := os.Open(*filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lineChan := make(chan string)
	go tailLine(file, lineChan)

	for {
		line := <-lineChan
		println(line)

		if r.MatchString(line) == true {
			regexGroup := r.FindStringSubmatch(line)
			for k,v := range regexGroup {
				println("Match: [",k,"]=", v)
			}
		}
	}
}

func tailLine(f *os.File, lineChan chan string) {
	//go to the end of the log file
	f.Seek(0, 2)

	reader := bufio.NewReader(f)
	var buffer bytes.Buffer
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				buffer.WriteString(line)
			} else {
				log.Panic(err)
			}
		} else {
			if buffer.Len() > 0 {
				buffer.WriteString(line)
				lineChan <- buffer.String()
				buffer.Reset()
			} else {
				lineChan <- line
			}

		}

		time.Sleep(1e7)
	}
}
