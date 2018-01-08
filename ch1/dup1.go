package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // Make creates an object of the given type
	// map[string]int (key: string | value: int)
	input := bufio.NewScanner(os.Stdin) // Reads the program standard input
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
