

package main

import (
	"fmt"
)

func main() {
	data := []int {1, 2, 3, 4, 5}
	fmt.Printf("Before: %v\n", data)
	fmt.Printf("Rotated by %d: %v\n", 2, rotateInOnePass(data, 2))
}

func rotateInOnePass(s []int, index int) []int {
	size := len(s)
	newArr := make([]int, size)
	for i, j := index, 0; j < size; i, j = i+1, j+1 {
		newArr[j] = s[i % size]
	}
	return newArr
}
