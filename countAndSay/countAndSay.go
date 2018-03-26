package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	n := []int{1, 2, 3, 4, 5}
	for _, m := range n {
		log.Printf("countAndSay: %v\n", countAndSay2(m))
	}
}

// 8ms, so fmt.Sprintf is very slow
func countAndSay1(n int) string {
	out := "1"
	for loop := 1; loop < n; loop++ {
		s, i, cuIndex := out, 0, 0
		out = ""
		for ; i < len(s); i++ {
			if s[cuIndex] != s[i] {
				out = out + fmt.Sprintf("%d%s", i-cuIndex, string(s[cuIndex]))
				cuIndex = i
			}
		}
		out = out + fmt.Sprintf("%d%s", i-cuIndex, string(s[cuIndex]))
	}
	return out
}

// lees than 1ms
func countAndSay2(n int) string {
	out := []string{"1"}
	for loop := 1; loop < n; loop++ {
		s, i, cuIndex := out, 0, 0
		out = []string{}
		for ; i < len(s); i++ {
			if s[cuIndex] != s[i] {
				out = append(out, strconv.Itoa(i-cuIndex), string(s[cuIndex]))
				cuIndex = i
			}
		}
		out = append(out, strconv.Itoa(i-cuIndex), string(s[cuIndex]))
	}
	return strings.Join(out, "")
}
