package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"unicode"
)

type counts struct {
	Lines      int
	Words      int
	Characters int
}

func (cs *counts) add(other counts) {
	cs.Lines += other.Lines
	cs.Words += other.Words
	cs.Characters += other.Characters
}

func countFile(fp string) (counts, error) {
	file, err := os.Open(fp)
	if err != nil {
		return counts{}, err
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
			return counts{}, err
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

	return counts{
		Lines:      lineCount,
		Words:      wordCount,
		Characters: characterCount,
	}, nil
}

func main() {
	filePaths := os.Args[1:]
	if len(filePaths) == 0 {
		fmt.Fprintln(os.Stderr, "usage: wordcount <file>")
		os.Exit(1)
	}
	hasFailed := false
	var total counts

	for _, filePath := range filePaths {
		c, err := countFile(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wordcount: %v\n", err)
			hasFailed = true
			continue
		}
		fmt.Printf("%d %d %d %s\n", c.Lines, c.Words, c.Characters, filePath)
		total.add(c)
	}
	if len(filePaths) > 1 {
		fmt.Printf("%d %d %d total\n", total.Lines, total.Words, total.Characters)
	}
	if hasFailed {
		os.Exit(1)
	}
}
