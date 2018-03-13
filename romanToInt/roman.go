package main

import (
	"fmt"
)

func main() {
	s := "DCX"
	n := romanToInt(s)
	fmt.Println(n)
}

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum, len := 0, len(s)
	for i := 0; i < len; i++ {
		val := m[s[i]]
		if i < len-1 && m[s[i]] < m[s[i+1]] {
			sum -= val
		} else {
			sum += val
		}
	}
	return sum
}
