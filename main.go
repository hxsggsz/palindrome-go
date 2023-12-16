package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	word := ""

	fmt.Print("A word to check if is palindrome: ")
	fmt.Scanln(&word)

	isPalindrome := checkPalindrome(word)

	if !isPalindrome {
		fmt.Println("It's not a palindrome")
		return
	}

	fmt.Println("It's a palindrome")
}

func checkPalindrome(word string) bool {
	splitedWord := strings.Split(word, "")
	slices.Reverse(splitedWord)
	joinedReverseString := strings.Join(splitedWord, "")

	return joinedReverseString == word
}
