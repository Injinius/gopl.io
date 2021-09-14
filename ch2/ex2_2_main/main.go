package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"gopl.io/ch2/ex2_2"
)

func main() {
	if len(os.Args) == 1 {
		consoleConvert()
	} else {
		convert(os.Args[1:])
	}
}

func consoleConvert() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Fprint(os.Stdout, "Enter numbers separated by spaces. Ctrl+C to exit.\n> ")
	for scanner.Scan() {
		convert(strings.Split(scanner.Text(), " "))
	}	
}


func convert(nums []string) {
	for _, arg := range nums {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatalf("main: %v\n", err)
		}
		showTemps(t)
		showWeights(t)
		//showLengths(t)
	}
}

func showTemps(temp float64) {
	f := ex2_2.Fahrenheit(temp)
	fmt.Printf("%v => %v, %v\n", f, ex2_2.FToC(f), ex2_2.FToK(f))
	c := ex2_2.Celsius(temp)
	fmt.Printf("%v => %v, %v\n", c, ex2_2.CToF(c), ex2_2.CToK(c))
	k := ex2_2.Kelvin(temp)
	fmt.Printf("%v => %v, %v\n", k, ex2_2.KToC(k), ex2_2.KToF(k))
}
func showWeights(weight float64) {
	p := ex2_2.Pound(weight)
	fmt.Printf("%v => %v\n", p, ex2_2.PoundsToKilos(p))
	k := ex2_2.Kilo(weight)
	fmt.Printf("%v => %v\n", k, ex2_2.KilosToPounds(k))
}