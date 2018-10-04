package main

import (
	"fmt"
)

//IsUnique -- This function will determine if the string passed in is made up of all unique characters.
func IsUnique(str string) bool {
	charMap := make(map[rune]bool)

	for index, runeValue := range str {
		fmt.Printf("character %v at %v \n", runeValue, index)
		_, ok := charMap[runeValue]
		if ok == true {
			return false
		}
		charMap[runeValue] = true
	}
	return true
}

//IsPermutation -- Compare two strings and determine if they are permutations of each other
func IsPermutation(str1, str2 string) bool {
	//I chose to use a map opposed to an array of 128 ints (ASCII). With this method we are not limited to the ASCII set and its probably more memory efficient as we need a map only as long as the string.

	if len(str1) != len(str2) {
		//cannot be a permutation if string are a different length
		return false
	}
	charMap1 := make(map[rune]int)

	for _, runeValue := range str1 {
		_, ok := charMap1[runeValue]
		if ok == true {
			charMap1[runeValue]++
		} else {
			charMap1[runeValue] = 1
		}
	}

	for _, runeValue := range str2 {
		count, ok := charMap1[runeValue]
		if ok == false {
			return false
		}
		count--
		if count < 0 {
			return false
		}
		charMap1[runeValue] = count
	}
	return true
}
