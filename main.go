package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const filePath = "messages.txt"

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("could not open %s: %s\n", filePath, err)
	}

	strChan := getLinesChannel(f)
	for line := range strChan {
		fmt.Printf("read: %s\n", line)
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	bytes := make([]byte, 8)
	var currentLine strings.Builder
	strChan := make(chan string)

	go func() {
		defer f.Close()
		defer close(strChan)
		for {
			n, err := f.Read(bytes)
			parts := strings.Split(string(bytes[:n]), string("\n"))

			for i := range parts {
				if i < len(parts)-1 {
					currentLine.WriteString(parts[i])
					strChan <- currentLine.String()
					currentLine.Reset()
				}
			}
			currentLine.WriteString(parts[len(parts)-1])

			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				} else {
					fmt.Println("Error reading file")
					break
				}
			}
		}
		strChan <- currentLine.String()
	}()
	return strChan
}
