package main

import (
	"log"
	"strings"
)

func main() {
	cases := [][]string{
		{"abcde", "bcdea"},
		{"abcde", "eabcd"},
		{"abcde", "acdeb"},
	}

	for _, s := range cases {
		log.Println(rotateString(s[0], s[1]))
	}
}

func rotateString(A string, B string) bool {
	return len(A) == len(B) && strings.Contains(A+A, B)
}
