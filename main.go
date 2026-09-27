package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"unicode"
)

func countFile(fp string) (int, int, int, error) {
	file, err := os.Open(fp)
	if err != nil {
		return 0, 0, 0, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	lineCount := 0
	wordCount := 0
	characterCount := 0
	previousWasSpace := true
	for {
		r, _, err := reader.ReadRune()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return 0, 0, 0, err
		}
		characterCount++
		if r == '\n' {
			lineCount++
		}
		if !unicode.IsSpace(r) && previousWasSpace {
			wordCount++
		}
		previousWasSpace = unicode.IsSpace(r)
	}

	return lineCount, wordCount, characterCount, nil
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Fprintln(os.Stderr, "usage: wordcount <file>")
		os.Exit(1)
	}

	for _, arg := range args[1:] {
		lineCount, wordCount, characterCount, err := countFile(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wordcount: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%d %d %d %s\n", lineCount, wordCount, characterCount, arg)
	}
}
