package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Fprintln(os.Stderr, "usage: wordcount <file>")
		os.Exit(1)
	}

	for _, arg := range args[1:] {
		fmt.Println(arg)
	}
}
