package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := " ", " "
	for _, arg := range os.Args[1:] { // Range iterates over elements in a variety of data structures.
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
