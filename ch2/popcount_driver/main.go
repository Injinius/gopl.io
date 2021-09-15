package main

import (
	"fmt"
	"gopl.io/ch2/popcount"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	/*
	pc=>[0 1 1 2 1 2 2 3 1 2 2 3 2 3 
	     3 4 1 2 2 3 2 3 3 4 2 3 3 4 
		 3 4 4 5 1 2 2 3 2 3 3 4 2 3 
		 3 4 3 4 4 5 2 3 3 4 3 4 4 5 
		 3 4 4 5 4 5 5 6 1 2 2 3 2 3 
		 3 4 2 3 3 4 3 4 4 5 2 3 3 4 
		 3 4 4 5 3 4 4 5 4 5 5 6 2 3 
		 3 4 3 4 4 5 3 4 4 5 4 5 5 6 
		 3 4 4 5 4 5 5 6 4 5 5 6 5 6 
		 6 7 1 2 2 3 2 3 3 4 2 3 3 4 
		 3 4 4 5 2 3 3 4 3 4 4 5 3 4 
		 4 5 4 5 5 6 2 3 3 4 3 4 4 5 
		 3 4 4 5 4 5 5 6 3 4 4 5 4 5 
		 5 6 4 5 5 6 5 6 6 7 2 3 3 4 
		 3 4 4 5 3 4 4 5 4 5 5 6 3 4 
		 4 5 4 5 5 6 4 5 5 6 5 6 6 7 
		 3 4 4 5 4 5 5 6 4 5 5 6 5 6 
		 6 7 4 5 5 6 5 6 6 7 5 6 6 7 
		 6 7 7 8]
	*/
}

func main() {
	for i, val := range pc {
		fmt.Printf("pc[%d]: %d => %b\n", i, val, i)
	}
	// So pc[204]: 4 => 11001100 --> pc[12]/1100 + pc[12]/1100 = 4
	fmt.Printf("Popcount for 204 (should be 4): %d\n", popcount.PopCount(204))
}
