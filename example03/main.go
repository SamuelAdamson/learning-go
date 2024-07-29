// Simple find and replace text using built-in string operations
// in Go.

// Run with `go run main.go [text to find] [replacement text] [body to search thru and find/replace]`

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check for number of arguments
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Not enough arguments (3 required).")
		os.Exit(-1)
	}

	// Get text to find and replacement
	find := os.Args[1]
	replace := os.Args[2]

	// Create scanner to read body text
	sc := bufio.NewScanner(os.Stdin)

	// Read through text, find and replace (print result to Stdout)
	for sc.Scan() {
		split := strings.Split(sc.Text(), find)
		joined := strings.Join(split, replace)

		fmt.Println(joined)
	}

}
