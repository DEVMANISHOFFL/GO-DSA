package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !unicode.IsLetter(rune(s[l])) && !unicode.IsDigit(rune(s[l])) {
			l++
		}

		for l < r && !unicode.IsLetter(rune(s[r])) && !unicode.IsDigit(rune(s[r])) {
			r--
		}

		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false
	fmt.Println(isPalindrome(" "))                              // true
}
