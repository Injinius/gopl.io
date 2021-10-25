// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5} // array not slice
	fmt.Println("Before: ", a) // "[5 4 3 2 1 0]"
	reverse(&a)  // Use pointer to array (not a slice)
	fmt.Println("Reversed: ", a) // "[5 4 3 2 1 0]"
	//!-array

}

//!+rev
// reverse reverses a slice of ints in place.
// Since a slice contains a pointer to an element of an array, passing a slice to a function per mits
// the function to modif y the underlying array elements

func reverse(s *[6]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}


//!-rev
