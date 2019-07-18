package shorterPath

// package main

import "fmt"

func main() {
	mmap := map[string]int{}
	fmt.Println("len(0)", len(mmap))

	mmap["a"] = 1
	fmt.Println("len(1)", len(mmap))

	mmap["b"] = 1
	fmt.Println("len(2)", len(mmap))
}
