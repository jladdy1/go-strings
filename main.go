package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello strings \n")

	isUnique := IsUnique("abcdefghi")
	fmt.Printf("IsUnique %v \n", isUnique)

	isPermutation := IsPermutation("dog", "god")
	fmt.Printf("IsPermutation (true) %v \n", isPermutation)
}
