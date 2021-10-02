package main

import (
	"fmt"
)

func main() {
	var a [3]int
	fmt.Printf("%T\n", a)
	fmt.Println("a[0] =", a[0])
	fmt.Println("a[len(a)-1] =", a[len(a)-1])

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	var q = [3]int{1, 2, 3}
	fmt.Println(q[2])
	var r = [3]int{1, 2} // Note: [3] sized, but only first two elements are initialized
	fmt.Println(r[2])
	s := [...]int{1, 2, 3} //  Note: size implied by initializer
	fmt.Printf("%T\n", s)


	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD:"$", EUR:"€", GBP:"₤", RMB:"¥"} //Initializing with indices
	fmt.Println("\n", EUR, symbol[EUR]) 

	z := [...]int{99: -1} // Allocates 100 element array 0-99
	fmt.Printf("\ntype=%T, z[0]=%d (default), z[99]=%d (explicit)\n", z, z[0], z[99]) 

	t := [2]int{1, 2}
	u := [...]int{1, 2}
	v := [2]int{1,3}
	fmt.Printf("\nt==u: %t, u==v: %t, t==v: %t\n", t==u, t==v, u==v)

// 	w := [3]int {1, 2}
// 	fmt.Println(t==w) //cannot compare t == w (mismatched types [2]int and [3]int)
}