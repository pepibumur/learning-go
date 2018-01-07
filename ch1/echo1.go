package main

import (
	"fmt" // Implements formatted I/O with functions
	"os"  // Platform-independent interface to operating system functionality
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
