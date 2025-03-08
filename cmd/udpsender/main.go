package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	add, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatalf("Unable to Resolve")
	}
	conn, err := net.DialUDP("udp", nil, add)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Cannot Read Stdin")
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Fatalf("Unable to Write to UDP")
		}
	}

}
