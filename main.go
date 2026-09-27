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
	args := os.Args
	if len(args) <= 1 {
		fmt.Fprintln(os.Stderr, "usage: wordcount <file>")
		os.Exit(1)
	}

	for _, arg := range args[1:] {
		c, err := countFile(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wordcount: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%d %d %d %s\n", c.Lines, c.Words, c.Characters, arg)
	}
}
