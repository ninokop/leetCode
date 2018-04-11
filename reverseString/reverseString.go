package main

import (
	"log"
)

func main() {
	cases := []string{
		"abcde",
		"reverseString",
		"hello world",
	}
	for _, s := range cases {
		log.Println(reverseString(s))
		log.Println(leftRotateString(s, 3))
	}
}

func reverseString(s string) string {
	sr := []rune(s)
	for i, j := 0, len(sr)-1; i <= j; i++ {
		sr[i], sr[j] = sr[j], sr[i]
		j--
	}
	return string(sr)
}

func leftRotateString(str string, left int) string {
	n := len(str)
	m := left % n

	strL := reverseString(string(str[0:m]))
	strR := reverseString(string(str[m:n]))
	return reverseString(string((strL + strR)[0:n]))
}
