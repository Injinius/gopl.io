// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 91.


// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

func main() {
	//!+main
	data := []string{"one", "", "three"}
	fmt.Println("nonempty version")
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`
	data = []string{"one", "", "three"}
	fmt.Println("nonempty2 version")
	fmt.Printf("%q\n", nonempty2(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`
	//!-main
}

//!+nonempty

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

//!-nonempty

//!+alt
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

//!-alt
