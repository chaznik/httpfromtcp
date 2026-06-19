package main

import (
	"fmt"
	"log"
	"os"
	"io"
	"bytes"
	"errors"
)

const filePath = "../messages.txt"

func main() {
	fd, err := os.Open(filePath)

	if err != nil {
		log.Fatalf("Error opening file: %s %v\n", filePath, err)
	}
	defer fd.Close()

	bytesRead := make([]byte, 8, 8)
	currentLine := ""

	for {
		n, err := fd.Read(bytesRead)

		if err != nil {
			if currentLine != "" {
				fmt.Printf("read: %s\n", currentLine)
				currentLine = ""
			}
			if errors.Is(err, io.EOF) {
				break;
			}
			fmt.Printf("error: %s\n", err.Error())
			break
		}

		parts := bytes.Split(bytesRead[:n], []byte("\n"))

		for i:=0; i < len(parts) - 1; i++ {
			fmt.Printf("read: %s%s\n", currentLine, parts[i])
			currentLine = ""
		}
		currentLine += string(parts[len(parts)-1])
	}
}
