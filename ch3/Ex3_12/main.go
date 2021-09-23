package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Example: main string1 string2")
	}
	string1, string2 := os.Args[1], os.Args[2]
	fmt.Printf("%s and %s %s anagrams.\n", string1, string2, areAnagrams(string1, string2))
}

// Exercise 3.12: Write a function that reports whether two strings are anagrams of each other,
// that is, they contain the same letters in a different order.
func areAnagrams(s1, s2 string) string {
	const ARE = "are"
	const ARENOT = ARE + " not"

	if len(s1) != len(s2) {
		return ARENOT
	}

	counts := make(map[byte]int)
	b1 := []byte(s1)
	b2 := []byte(s2)
	for i:=0; i < len(s1); i++ {
		counts[b1[i]]++
		counts[b2[i]]--
	}
	for _, count := range counts {
      if count != 0 {
		  return ARENOT
	  }
	}
	return ARE
}