package main

import "fmt"

func main() {
	// Sprintf formats according to a format specifier and returns the resulting string.
	words := fmt.Sprintf("Hello World %s", "aasumitro")
	// Printf formats according to a format specifier and writes to standard output.
	fmt.Println(words)
	// words is a Short variable declarations
	// for more https://tour.golang.org/basics/10
}

