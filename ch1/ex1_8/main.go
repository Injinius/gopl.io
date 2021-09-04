// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)
const (
	HTTP_PREFIX = "http://"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, HTTP_PREFIX) {
			url = fmt.Sprintf("%s%s", HTTP_PREFIX, url)
			fmt.Printf("ex1_8: Re-writing url as %s\n", url)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex1_8: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex1_8: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("\nex1_8: Bytes written: %d\n", b)
	}
}

//!-
