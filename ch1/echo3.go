package main

import (
	"fmt"
	"os"
	"strings" // Package strings implements simple functions to manipulate UTF-8 encoded strings.
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
