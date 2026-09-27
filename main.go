package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

// countReader counts lines, words, and characters the same way wc -lwm does:
// a line is a '\n', a word is a run of non-whitespace, a character is a rune.
func countReader(r io.Reader) (counts, error) {
	reader := bufio.NewReader(r)
	var c counts
	previousWasSpace := true
	for {
		ch, _, err := reader.ReadRune()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return counts{}, err
		}
		c.Characters++
		if ch == '\n' {
			c.Lines++
		}
		isSpace := unicode.IsSpace(ch)
		if !isSpace && previousWasSpace {
			c.Words++
		}
		previousWasSpace = isSpace
	}
	return c, nil
}

func countFile(path string) (counts, error) {
	file, err := os.Open(path)
	if err != nil {
		return counts{}, err
	}
	defer file.Close()
	return countReader(file)
}

// formatCounts builds one output line with only the selected columns.
// An empty name is left off, which is what wc does for stdin.
func formatCounts(c counts, name string, lines, words, characters bool) string {
	var fields []string
	if lines {
		fields = append(fields, strconv.Itoa(c.Lines))
	}
	if words {
		fields = append(fields, strconv.Itoa(c.Words))
	}
	if characters {
		fields = append(fields, strconv.Itoa(c.Characters))
	}
	if name != "" {
		fields = append(fields, name)
	}
	return strings.Join(fields, " ")
}

func main() {
	lines := flag.Bool("l", false, "count lines")
	words := flag.Bool("w", false, "count words")
	characters := flag.Bool("m", false, "count characters")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage: wordcount [-l] [-w] [-m] [file ...]")
		flag.PrintDefaults()
	}
	flag.Parse()

	if !*lines && !*words && !*characters {
		*lines = true
		*words = true
		*characters = true
	}

	filePaths := flag.Args()
	if len(filePaths) == 0 {
		c, err := countReader(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wordcount: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(formatCounts(c, "", *lines, *words, *characters))
		return
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
		fmt.Println(formatCounts(c, filePath, *lines, *words, *characters))
		total.add(c)
	}

	if len(filePaths) > 1 {
		fmt.Println(formatCounts(total, "total", *lines, *words, *characters))
	}

	if hasFailed {
		os.Exit(1)
	}
}
