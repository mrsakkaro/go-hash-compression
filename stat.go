package main

import (
	"os"
	"fmt"
	_ "io"
	"path/filepath"
	_ "time"
)

func main() {

	src, err := os.Open("compress-me.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := src.Close(); err != nil {
			panic(err)
		}
	}()


	res, err := os.Stat("golang.tar.gz")
	lstat, err := os.Lstat("golang.tar.gz")

	fmt.Printf("\n%s\n", lstat)

	fmt.Printf("\n%s\n", res)
	/*
	fmt.Printf("\n%s\n", res.Name())
	fmt.Printf("%s\n", res.Size())
	fmt.Printf("%s\n", res.ModTime().String())
	*/

	dir, err := filepath.Abs("golang.tar.gz")
	fmt.Printf("\n%s\n", dir)
}
