package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// work on files by related packages
	// read subs directory
	files, _ := os.ReadDir("subs")
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".srt") {
			// show files name ... just Only with extensions .srt
			fmt.Println(file.Name())
		}
	}
}
