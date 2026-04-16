package main

import (
	"fmt"
	"os"
	"log"
	"io"
)

const filePath = "messages.txt"

func main() {
	fd, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error - %x opening file: %s", err, filePath)
	}
	defer fd.Close()

	data := make([]byte, 8)

	for {
		n, err := fd.Read(data) 
		if err != nil {
			log.Printf("Error reading file: %s", err) 
		}

		if n > 0 {
			fmt.Printf("read: %s\n", data[:n]) 
		}

		if err == io.EOF {
			break
		}
	}
}
