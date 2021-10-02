package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"log"
	"os"
)

const (
	SHA256 = "SHA256"
	SHA384 = "SHA384"
	SHA512 = "SHA512"
)

var hashPtr = flag.String("hash", SHA256, SHA256 + ", " + SHA384 + ", or " + SHA512)

func main() {
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	flag.Parse()
	fmt.Printf("Hash type: %s\n", *hashPtr)

	scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Enter text to be hashed: ")
        // Scans a line from Stdin(Console)
        scanner.Scan()
        // Holds the string that scanned
        text := scanner.Text()
		if (len(text) > 0) {
			hash, err := getHash([]byte(text))
			if err != nil {
				log.Fatal("Unexpected error. ", err)
			}
			fmt.Printf("Hash of %s: %x\n", text, hash)
		} else {
			fmt.Println("No text provided. Exiting.")
			break
		}
	}
}

func getHash(x []byte) ([]byte, error) {
	var res []byte
	var h hash.Hash
	var err error
 	switch *hashPtr {
	case SHA256:
		h = sha256.New()
	case SHA384:
		h = sha512.New384()
	case SHA512:
		h = sha512.New()
	default:
		err = fmt.Errorf("hash type %s not supported ", *hashPtr)
	}
	if h != nil {
		h.Write(x)
		res = h.Sum(nil)
	}
	return res, err
}