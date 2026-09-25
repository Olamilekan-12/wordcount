package main

import (
	"fmt"
	"os"
)

func countFile(fp string) error {
	file, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Fprintln(os.Stderr, "usage: wordcount <file>")
		os.Exit(1)
	}

	for _, arg := range args[1:] {
		fmt.Println(arg)
		err := countFile(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wordcount: %v\n", err)
			os.Exit(1)
		}
	}
}
