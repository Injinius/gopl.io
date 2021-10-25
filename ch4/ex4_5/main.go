

package main

import (
	"fmt"
)
//Exercise 4.5: 
// Write an in-place function to eliminate adjacent 
// duplicates in a []string slice.
func main() {
	// two pointers: read -> reads through original data, write -> increments for unique values
	// eliminating duplicates means the original array will always have capacity
	//data := []string {}
	//data := []string {"a"}
	data := []string {"a", "a", "a", "b", "c", "d", "d"}
	fmt.Printf("Before: %v, size: %d\n", data, len(data))
	data = deDup(data) // in-place modification, but you need to return a potentially modified slice
	fmt.Printf("After:  %v, size: %d\n", data, len(data))
}

func deDup(data []string) []string {
	size := len(data)
	if (size == 0) {
		return data
	}

	write := 1
	for read := 1; read < size; read = read + 1 {
		if data[read] != data[read-1] {
			data[write] = data[read]
			write = write + 1
		}
	}
	return data[:write]
}