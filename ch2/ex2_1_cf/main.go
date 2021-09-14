// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/ex2_1"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := ex2_1.Fahrenheit(t)
		c := ex2_1.Celsius(t)
		k := ex2_1.Kelvin(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, ex2_1.FToC(f), f, ex2_1.FToK(f))
		fmt.Printf("%s = %s, %s = %s\n",
			c, ex2_1.CToF(c), c, ex2_1.CToK(c))
		fmt.Printf("%s = %s, %s = %s\n",
			k, ex2_1.KToF(k), k, ex2_1.KToC(k))
	}
}

//!-
