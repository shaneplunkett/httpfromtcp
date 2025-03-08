package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const message_path = "./messages.txt"

func openFile(f string) (*os.File, error) {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	return file, nil
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		defer f.Close()
		data := make([]byte, 8)
		var currentLine string
		for {
			count, err := f.Read(data)
			if err != nil {
				if err == io.EOF {
					if currentLine != "" {
						ch <- currentLine
					}
					break
				}
				break
			}
			part := string(data[:count])
			parts := strings.Split(part, "\n")
			for _, part := range parts[:len(parts)-1] {
				currentLine += part
				ch <- currentLine
				currentLine = ""
			}
			if len(parts) > 0 {
				currentLine += parts[len(parts)-1]
			}
		}
	}()
	return ch
}

func main() {
	f, err := openFile(message_path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	ch := getLinesChannel(f)
	for l := range ch {
		fmt.Printf("read: %s\n", l)
	}

}
