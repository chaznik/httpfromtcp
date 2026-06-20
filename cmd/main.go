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

	lines := getLinesChannel(fd)

	for line:= range lines {
		fmt.Printf("read: %s\n", line)
	}
}


func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)
	go func() {
		defer fd.Close()
		defer close(lines)
		bytesRead := make([]byte, 8)
		currentLine := ""

		for {
			n, err := fd.Read(bytesRead)

			if err != nil {
				if currentLine != "" {
					lines <- currentLine
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
				lines <- fmt.SPrintf("%s%s", currentLine, parts[i])
				currentLine = ""
			}
			currentLine += string(parts[len(parts)-1])
		}
	}()
	return lines
}
