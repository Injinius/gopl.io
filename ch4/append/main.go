// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 88.

// Append illustrates the behavior of the built-in append function.
package main

import "fmt"

func appendslice(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to expand the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

//!+append
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	//fmt.Printf("x.len=%d, x.cap=%d, zlen=%d\n", len(x), cap(x), zlen)
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		//fmt.Printf("zcap=%d, len*2=%d\n", zcap, len(x)*2)
		if zcap < 2*len(x) { //Note: When len=0,1 -- there's no bump
			zcap = 2 * len(x)
			//fmt.Printf("bumped zcap=%d\n", zcap)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

//!-append

//!+growth
func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
	x = appendslice(x, x...)
	fmt.Printf("appendslice: cap=%d\t%v\n", cap(x), x)
	x = appendslice(x, 100, 101, 102)
	fmt.Printf("appendslice: cap=%d\t%v\n", cap(x), x)

}

//!-growth

/*
//!+output
0  cap=1   [0]
1  cap=2   [0 1]
2  cap=4   [0 1 2]
3  cap=4   [0 1 2 3]
4  cap=8   [0 1 2 3 4]
5  cap=8   [0 1 2 3 4 5]
6  cap=8   [0 1 2 3 4 5 6]
7  cap=8   [0 1 2 3 4 5 6 7]
8  cap=16  [0 1 2 3 4 5 6 7 8]
9  cap=16  [0 1 2 3 4 5 6 7 8 9]
//!-output
*/
