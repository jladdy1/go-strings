package main

import "fmt"

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
