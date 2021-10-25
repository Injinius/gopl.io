package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("Data=%v\n", s)
	rotate(s, 2) // in-place rotation
	fmt.Printf("Rotated by %d: %v\n", 2, s)
}

func rotate(s []int, i int) {
	reverse(s[:2])
    reverse(s[2:])
    reverse(s)
}

// "ch4/rev"
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}