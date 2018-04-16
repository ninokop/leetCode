package main

import (
	"log"
)

func main() {
	cases := []string{
		"babad",
		"cbbd",
		"allolll",
	}

	for _, s := range cases {
		log.Println(longestPalindrome(s))
	}
}

func longestPalindrome2Way(s string) string {
	var (
		out          string
		max, c, i, j int
	)

	n := len(s)
	for i = 0; i < n; i++ {
		for j = 0; (i-j >= 0) && (i+j < n); j++ {
			if s[i-j] != s[i+j] {
				break
			}
			c = 2*j + 1
		}
		if c > max {
			max = c
			out = s[i-j+1 : i+j]
		}
	}

	for i = 0; i < n; i++ {
		for j = 0; (i-j >= 0) && (i+j+1 < n); j++ {
			if s[i-j] != s[i+j+1] {
				break
			}
			c = 2*j + 2
		}
		if c > max {
			max = c
			out = s[i-j+1 : i+j+1]
		}
	}
	return out
}

func longestPalindrome(s string) string {
	var (
		pos, maxR, maxLen int
		out               []byte

		n  = len(s)
		RL = make([]int, 2*n+1)
		m  = make([]byte, 2*n+1)
	)

	for i := 0; i < n; i++ {
		m[i*2+1] = s[i]
		m[i*2] = '#'
	}
	m[2*n] = '#'

	for i := 0; i < 2*n+1; i++ {
		if i < maxR {
			RL[i] = min(RL[2*pos-i], maxR-i)
		} else {
			RL[i] = 1
		}

		for (i-RL[i] >= 0) && (i+RL[i] < 2*n+1) && m[i-RL[i]] == m[i+RL[i]] {
			RL[i]++
		}

		if RL[i]+i-1 > maxR {
			maxR = RL[i] + i - 1
			pos = i
		}

		if maxLen < RL[i] {
			maxLen = RL[i]
			out = m[i-RL[i]+1 : i+RL[i]]
		}

	}
	return trim(out)
}

func trim(s []byte) (out string) {
	for i := range s {
		if s[i] != '#' {
			out += string(s[i])
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
