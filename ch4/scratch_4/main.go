package main

import (
	"fmt"
)

func main() {
	a := make([]string, 10)
	p(a)
	b := make([]string, 10, 15)
	p(b)

	var x []int
	x = append(x, 1) // append might change underlying array, so you need to (re-)assign
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) // append the slice x
	fmt.Println("Appended slice:", x) // "[1 2 3 4 5 6 1 2 3 4 5 6]"
}

func p(x []string)  {
	fmt.Printf("Data/len=%v/len=%d/cap=%d\n", x, len(x), cap(x))
}