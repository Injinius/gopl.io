// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package ex2_2

// PoundsToKilos converts pounds to kilos
func PoundsToKilos(p Pound) Kilo { return Kilo(p*Pound(pToK)) }

// KilosToPounds converts kilos to pounds
func KilosToPounds(k Kilo) Pound { return Pound(k*Kilo(kToP)) }

//!-
