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
	fd, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("could not open %s: %s\n", filePath, err)
	}
	defer fd.Close()

	bytes := make([]byte, 8)
	currentLine := ""

	for {
		n, err := fd.Read(bytes)
		parts := strings.Split(string(bytes[:n]), string("\n"))

		for i := range parts {
			if i < len(parts)-1 {
				currentLine += parts[i]
				fmt.Printf("read: %s\n", currentLine)
				currentLine = ""
			}
		}
		currentLine += parts[len(parts)-1]

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				fmt.Println("Error reading file")
			}
		}
	}

	fmt.Printf("read: %s\n", currentLine)
}
