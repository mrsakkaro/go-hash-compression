package main

import (
	"crypto/sha256"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"io"
)

func main() {

	// Open Source File
	src, err := os.Open("compress-me.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := src.Close(); err != nil {
			panic(err)
		}
	}()

	h2 := sha256.New()
	h1 := sha1.New()
	h5 := md5.New()

	if _, err := io.Copy(h5, src); err != nil {
		panic(err)
	}
	if _, err := io.Copy(h2, src); err != nil {
		panic(err)
	}
	if _, err := io.Copy(h1, src); err != nil {
		panic(err)
	}

	fmt.Printf("\n%s\n", src)
	fmt.Printf("\n%s\n", src)
	fmt.Printf("\n%s\n", src)

	fmt.Printf("Sha256:\n%x\n", h2.Sum(nil))
	fmt.Printf("Sha1:\n%x\n", h1.Sum(nil))
	fmt.Printf("Sha1:\n%x\n", hex.EncodeToString(h1.Sum(nil)))
	fmt.Printf("Sha1:\n%x\n", h1.Sum(nil)[:20])
	fmt.Printf("Sha1:\n%x\n", hex.EncodeToString(h1.Sum(nil)[:20]))
	fmt.Printf("md5:\n%d\n", h5.Sum(nil))
	fmt.Printf("md5:\n%d\n", hex.EncodeToString(h5.Sum(nil)))
	fmt.Printf("md5:\n%d\n", h5.Sum(nil)[:16])
	fmt.Printf("md5:\n%d\n", hex.EncodeToString(h5.Sum(nil)[:16]))
}
