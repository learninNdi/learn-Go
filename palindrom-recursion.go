package main

import "fmt"

func isPalindromCheck(value string, idx int) bool {
	if idx >= len(value)/2 {
		return true
	}

	if value[idx] != value[len(value)-1-idx] {
		return false
	}

	return isPalindromCheck(value, idx+1)
}

func isPalindrom(value string) bool {
	return isPalindromCheck(value, 0)
}

func main() {
	fmt.Println(isPalindrom(""))      // true
	fmt.Println(isPalindrom("aku"))   // false
	fmt.Println(isPalindrom("aka"))   // true
	fmt.Println(isPalindrom("kodok")) // true
	fmt.Println(isPalindrom("katak")) // true
	fmt.Println(isPalindrom("dia"))   // false
	fmt.Println(isPalindrom("ada"))   // true
}
