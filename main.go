package main

import (
	"os"
	"path/filepath"
	"strings"
)

const subs_dir = "subs"

func main() {
	// work on files by related packages
	// read subs directory
	files, _ := os.ReadDir("subs_dir")
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".srt") || strings.HasSuffix(filename, ".vtt") {
			// show files name ... just Only with extensions .srt
			Process(filename)

		}
	}
}

// process file
func Process(filename string) {

	// read file
	path := filepath.Join(sub_dir, filename)
	bytes, _ := os.ReadFile(path)
	content := string(bytes)

	// divide files to lines
	lines := strings.Split(content, "\n")

}

// process lines
// for finding just string
// Algorithm for deleting anything except strings
func ExtractWords(lines []string) []string {
	words := []string{}
	indexer := 0
	for _, line := range lines {
		if line == "\r" {
			indexer += 2
			continue
			// "r" is a method in subtitles ... is difference between two writed conversation
		}
		if indexer > 0 {
			indexer--
			continue
		}
		new_words := strings.Split(line, " ")
		// the line that we want (there is words in it)
		// add new words to words
		words = append(words, new_words...)
	}
	return words
}
