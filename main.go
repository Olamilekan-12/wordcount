package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countFile(fp string) (int, int, error) {
	file, err := os.Open(fp)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineCount := 0
	wordCount := 0
	for scanner.Scan() {
		wordCount += len(strings.Fields(scanner.Text()))
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		return 0, 0, err
	}
	return lineCount, wordCount, nil
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Fprintln(os.Stderr, "usage: wordcount <file>")
		os.Exit(1)
	}

	for _, arg := range args[1:] {
		lineCount, wordCount, err := countFile(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wordcount: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%d %d %s\n", lineCount, wordCount, arg)
	}
}
