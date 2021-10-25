package main

import (
	"fmt"
)

func main() {
	months := [...]string{
		1: "January", 
		2: "February", 
		3: "March", 
		4: "April", 
		5: "May", 
		6: "June", 
		7: "July", 
		8: "August", 
		9: "September", 
		10: "October", 
		11: "Noveember", 
		12: "December"}
		Q2 := months[4:7]
		summer := months[6:9]
		fmt.Println(Q2) // ["April" "May" "June"]
		fmt.Println(summer) // ["June" "July" "August"]
		for _, s := range summer { //(inefficient) test for common elements
			for _, q := range Q2 {
				if s == q {
					fmt.Printf("%s appears in both\n", s)
				}
			}
		}
		
		// Slicing beyond cap(s) causes a panic, 
		// fmt.Println(summer[:20]) // panic: out of range

		// but slicing beyond len(s) extends the slice, so the result may be longer than the original:
		endlessSummer := summer[:5] // extend a slice (within capacity)
		fmt.Println(endlessSummer) // "[June July August September October]"
	}