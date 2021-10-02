package main

import (
	"fmt"
	"crypto/sha256"
)

func main() {
	three := 0b0011
	five :=  0b0101
	fmt.Printf("3:     %04b\n", three)
	fmt.Printf("5:     %04b\n", five)
	fmt.Printf("3 & 5: %04b\n", three & five)
	fmt.Printf("3 ^ 5: %04b\n", three ^ five)
	// 3:     0011
	// 5:     0101
	// 3 & 5: 0001
	// 3 ^ 5: 0110 <-- # of different bits
	b3 := [32]byte {0:3, 31:3}
	b5 := [32]byte {0:5, 31:5}
	fmt.Println("Popcount(3^5): ", popCount(b3, b5))
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println("Popcount(c1^c2): ", popCount(c1, c2))
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popCount(x, y [32]byte) int {
	count := 0
	for i := range x {
		num := int(pc[x[i] ^ y[i]])
		count += num
		// fmt.Printf("%08b ^ %08b: %08b [%d][%3d]\n", x[i], y[i], x[i] ^ y[i], num, count)
	}
	return count
}