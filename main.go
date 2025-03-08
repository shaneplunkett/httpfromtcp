package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

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
	ln, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	fmt.Println("Setting up Server...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection Accepted...")
		ch := getLinesChannel(conn)
		for l := range ch {
			fmt.Printf("%s\n", l)
		}
		fmt.Println("Connection Closed...")

	}
}
