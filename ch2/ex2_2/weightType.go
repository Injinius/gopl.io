// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
package ex2_2

import "fmt"

type Pound float64
type Kilo float64

const (
	pToK float64 = 0.453592
	kToP float64 = 1.0 / pToK
)

func (p Pound) String() string    { return fmt.Sprintf("%.2f lbs", p) }
func (k Kilo) String() string { return fmt.Sprintf("%.2f kg", k) }

//!-
