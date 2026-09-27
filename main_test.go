package main

import (
	"os"
	"strings"
	"testing"
)

func TestCountReader(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  counts
	}{
		{"empty", "", counts{}},
		{"single line", "hello world\n", counts{Lines: 1, Words: 2, Characters: 12}},
		{"no trailing newline", "hello", counts{Lines: 0, Words: 1, Characters: 5}},
		{"multiple spaces", "hi   there\n", counts{Lines: 1, Words: 2, Characters: 11}},
		{"tabs and newlines", "a\tb\nc\n", counts{Lines: 2, Words: 3, Characters: 6}},
		{"only whitespace", " \n\t\n", counts{Lines: 2, Words: 0, Characters: 4}},
		{"unicode", "héllo 🚀", counts{Lines: 0, Words: 2, Characters: 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := countReader(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("countReader(%q) = %+v, want %+v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCountFile(t *testing.T) {
	got, err := countFile("testdata/unicode.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := counts{Lines: 0, Words: 2, Characters: 7}
	if got != want {
		t.Errorf("countFile = %+v, want %+v", got, want)
	}
}

func TestCountFileMissing(t *testing.T) {
	_, err := countFile("testdata/does-not-exist.txt")
	if !os.IsNotExist(err) {
		t.Errorf("expected a not-exist error, got %v", err)
	}
}

func TestCountsAdd(t *testing.T) {
	total := counts{Lines: 1, Words: 2, Characters: 3}
	total.add(counts{Lines: 10, Words: 20, Characters: 30})
	want := counts{Lines: 11, Words: 22, Characters: 33}
	if total != want {
		t.Errorf("add = %+v, want %+v", total, want)
	}
}

func TestFormatCounts(t *testing.T) {
	c := counts{Lines: 1, Words: 2, Characters: 3}
	tests := []struct {
		name                     string
		file                     string
		lines, words, characters bool
		want                     string
	}{
		{"all columns", "a.txt", true, true, true, "1 2 3 a.txt"},
		{"lines only", "a.txt", true, false, false, "1 a.txt"},
		{"words and characters", "a.txt", false, true, true, "2 3 a.txt"},
		{"stdin has no name", "", true, true, true, "1 2 3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatCounts(c, tt.file, tt.lines, tt.words, tt.characters)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
