package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", commaSigned(os.Args[i]))
	}
}

// Exercise 3.11: Enhance comma so that it deals correctly 
// with floating-point numbers and an optional sign.
func commaSigned(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	const ( 
		plus = "+"
		minus = "-"
		decimalPoint = "."
    )

	var sign string = ""
	var whole string = s

	switch s[0:1] {
	    case minus:
		    sign = minus
			whole = s[1:]
	    case plus:
		    sign = plus
			whole = s[1:]
	}

	fraction := ""
	fractionIndex := strings.Index(whole, decimalPoint)
	if (fractionIndex != -1) {
        fraction = whole[fractionIndex:]
		whole = whole[0: fractionIndex]
	}

	return sign + comma(whole) + fraction
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}