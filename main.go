package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func countWords(text string) int {
	return len(strings.Fields(text))
}

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println(`Usage: go run . "some text to count"`)
	// }

	text := flag.String("text", "", "text to count")
	file := flag.String("file", "", "path to text file")
	flag.Parse()

	// if strings.TrimSpace(*text) == "" {
	// 	fmt.Println("Error: provide text with -text")
	// 	flag.Usage()
	// 	return
	// }

	// text := strings.Join(os.Args[1:], " ")

	if *text != "" && *file != "" {
		fmt.Fprintln(os.Stderr, "Choose either -text or -file")
		os.Exit(1)
	}

	if strings.TrimSpace(*text) != "" {
		fmt.Println("Word count:", countWords(*text))
		return
	}

	if *file != "" {
		data, err := os.ReadFile(*file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading file:", err)
			os.Exit(1)
		}
		fmt.Println("Word count:", countWords(string(data)))
	}
}
