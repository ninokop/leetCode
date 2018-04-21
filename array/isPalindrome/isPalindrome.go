package main

import (
	"fmt"
)

func main() {
	x := 1234543214
	isP := isPalindrome(x)
	fmt.Println(isP)
}

func isPalindrome1(x int) bool {
	t, r := x, 0
	for t > 0 {
		r = r*10 + t%10
		t = t / 10
	}
	return x == r
}

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	rev := 0
	for rev < x {
		rev = rev*10 + x%10
		x = x / 10
	}
	return x == rev || x == rev/10
}
