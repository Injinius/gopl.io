// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	// "fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	var writer = os.Stdout
	if len(os.Args) > 1 { //Printing to Stdout on Windows and redirecting to a file produces invalif png content
		var err error
		writer, err = os.Create(os.Args[1])
		if err != nil {
			log.Fatalf("Could not open file=%s because %v\n", os.Args[1], err)
		}
	}

	dy := float64(ymax-ymin)/height
	dx := float64(xmax-xmin)/width
	var colors [width][height]color.Color
	for py := 0; py < height; py++ {
		y := float64(py)*dy + ymin
		for px := 0; px < width; px++ {
			x := float64(px)*dx + xmin
			colors[px][py] = mandelbrot(complex(x, y))
		}
	}
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// Compute pixel from avg of desired pixel with right, top, and upper right neighbors
	for py := 0; py < height-1; py++ {
		for px := 0; px < width-1; px++ {
			// fmt.Printf("px=%d, py=%d\n", px, py)

			r1, g1, b1, a1 := colors[px][py].RGBA()
			// fmt.Printf("r1=%d, g1=%d, b1=%d, a1=%d\n", r1, g1, b1, a1)
			r2, g2, b2, a2 := colors[px+1][py].RGBA()
			// fmt.Printf("r2=%d, g2=%d, b2=%d, a2=%d\n", r2, g2, b2, a2)
			r3, g3, b3, a3 := colors[px][py+1].RGBA()
			// fmt.Printf("r3=%d, g3=%d, b3=%d, a3=%d\n", r3, g3, b3, a3)
			r4, g4, b4, a4 := colors[px+1][py+1].RGBA()
			// fmt.Printf("r4=%d, g4=%d, b4=%d, a4=%d\n", r4, g4, b4, a4)
			// fmt.Printf("(r1 + r2 + r3 + r4)/4=%d\n", (r1 + r2 + r3 + r4)/4)
			// fmt.Printf("(g1 + g2 + g3 + g4)/4=%d\n", (g1 + g2 + g3 + g4)/4)
			// fmt.Printf("(b1 + b2 + b3 + b4)/4=%d\n", (b1 + b2 + b3 + b4)/4)
			// fmt.Printf("(a1 + a2 + a3 + a4)/4=%d\n", (a1 + a2 + a3 + a4)/4)
			avgColor := color.RGBA{
				uint8((r1 + r2 + r3 + r4)/1024),
				uint8((g1 + g2 + g3 + g4)/1024),
				uint8((b1 + b2 + b3 + b4)/1024),
				uint8((a1 + a2 + a3 + a4)/1024)}
			// fmt.Printf("avgColor=%v\n", avgColor)

			// Image point (px, py) represents complex value z.
			img.Set(px, py, avgColor)
		}
	}
	png.Encode(writer, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
