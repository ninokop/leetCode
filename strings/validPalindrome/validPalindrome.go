package main

import (
	"log"
	"strings"
)

func main() {
	cases := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		".,",
		"OP",
	}
	for _, s := range cases {
		log.Println(isPalindrome(s))
	}
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	s = strings.ToLower(s)
	n := len(s) - 1
	for i, j := 0, n; i < j; {
		for ; !isLetterOrDigit(s[i]) && i < j; i++ {
		}
		for ; !isLetterOrDigit(s[j]) && i < j; j-- {
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isLetterOrDigit(b byte) bool {
	if (b < 'a' || b > 'z') && (b < 'A' || b > 'Z') && (b < '0' || b > '9') {
		return false
	}
	return true
}
