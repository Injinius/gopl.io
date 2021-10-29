

package main

import (
	"fmt"
	"sort"
)

func main() {
	//ages := make(map[string]int) // mapping from strings to ints
	ages := map[string]int {
		"alice": 31,
		"charlie": 34,
		}
	ages["bob"] = 33 //Adding to map
	fmt.Println("Unsorted listing ...")
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	names := make([]string, 0, len(ages))
	for name := range ages {
        names = append(names, name)
    }
	sort.Strings(names)
	fmt.Println("Sorted listing ...")
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}



	// Checking for map existence
	delete(ages, "bob") // Ensure that bob isn't in the map
	age, ok := ages["bob"]
	if !ok { 
		fmt.Println("Bob is not in the map")
 	} else {
		fmt.Println("Bob is ", age)
	 }

	 //Or a more compact version
	 if age2, ok := ages["bob"]; !ok { fmt.Println("Bob is not in the map") } else { fmt.Println("Bob is ", age2) }

}