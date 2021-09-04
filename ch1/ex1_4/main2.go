// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Meta struct {
	Count int
	Files map[string]bool
}

func main() {
	counts := make(map[string]*Meta)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("Reading from Stdin")
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			fmt.Printf("Reading from %s\n", arg)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex1_4: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, md := range counts {
		if md.Count > 1 {

			fmt.Printf("%d\t[%v]\t%s\n", md.Count, md.Files, line)
		}
	}
}

func countLines(f *os.File, counts map[string]*Meta) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := input.Text()
		fmt.Printf("Line: %s\n",key)

		if counts[key] == nil {
			fmt.Printf("Init meta data\n")
			md := &Meta{Count: 0, Files: make(map[string]bool)}
			counts[key] = md
		}
		counts[key].Count++
		counts[key].Files[f.Name()] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}
//a
//a
//a

//!-
