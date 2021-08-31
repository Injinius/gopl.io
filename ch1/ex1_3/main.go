package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	fmt.Println("Echo1")
	benchmark(echo1)
	fmt.Println("Echo2")
	benchmark(echo2)
	fmt.Println("Echo3")
	benchmark(echo3)
}

func benchmark(f func()) {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		f()
	}
	fmt.Printf("%d ms elapsed\n", time.Since(start).Milliseconds())
}
func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
}
func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
}
func echo3() {
	strings.Join(os.Args[1:], " ")
}
