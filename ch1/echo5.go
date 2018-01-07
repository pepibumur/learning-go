package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	mainFor()
	mainJoin()
}

func mainFor() {
	timeBefore := time.Now()
	s, sep := " ", " "
	for _, arg := range os.Args[1:] { // Range iterates over elements in a variety of data structures.
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	timeAfter := time.Now()
	fmt.Printf("%v seconds using FOR\n", timeAfter.Sub(timeBefore))
}

func mainJoin() {
	timeBefore := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	timeAfter := time.Now()
	fmt.Printf("%v seconds using JOIN\n", timeAfter.Sub(timeBefore))
}
