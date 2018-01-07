package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Command: %s\n", os.Args[0])
	for index, value := range os.Args[1:] {
		fmt.Printf("Argument %d: %s\n", index, value)
	}
}
