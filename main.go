package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello strings /n")

	isUnique := IsUnique("abcdefghii")
	fmt.Printf("IsUnique %v", isUnique)
}
