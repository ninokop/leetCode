package main

import (
	"log"
	"unsafe"
)

func main() {
	cases := []string{
		"the sky is blue",
		"my name is nino",
		"Hello World",
	}

	for _, s := range cases {
		log.Println(reverseWords(s))
	}
}

func reverseWords(str string) string {
	s := StringToBytesWithNoCopy(str)
	var (
		out string
		cur int
	)
	for i := range s {
		if s[i] == ' ' {
			out += reverseString(string(s[cur:i])) + " "
			cur = i + 1
		}
	}

	out += reverseString(string(s[cur:len(s)]))
	return reverseString(out)
}

func reverseString(s string) string {
	sr := StringToBytesWithNoCopy(s)
	for i, j := 0, len(sr)-1; i <= j; i++ {
		sr[i], sr[j] = sr[j], sr[i]
		j--
	}
	return BytesToStringWithNoCopy(sr)
}

func BytesToStringWithNoCopy(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func StringToBytesWithNoCopy(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
