package main

import (
	"log"
)

func main() {
	cases := [][]string{
		{"HelloNino", "ello"},
		{"ABCDABDEFG", "ABDE"},
		{"hello", "lol"},
	}

	for _, s := range cases {
		log.Println(strStr(s[0], s[1]))
	}
}

func strStr(haystack string, needle string) int {
	nLen, hLen := len(needle), len(haystack)
	switch {
	case nLen == 0:
		return 0
	case nLen == hLen:
		if haystack == needle {
			return 0
		}
		return -1
	case nLen > hLen:
		return -1
	}
	return KMPSearch(haystack, needle)
}

/* How to get next
if next[j] = k, then s[0:k) = s[j-k:j), which includes s[k-1]=s[j-1]
then how to compute next[j+1]

if s[k] == s[j]; next[j+1] = k+1
if s[k] != s[j]; you need to find lagest index m (m<k)
which lead to s[0:m]=s[j-m:j]

since s[j-m:j-1]=s[k-m:k-1] is verified by next[j]=k
so s[0:m-1]=s[k-m:k-1], which means next[k] = m

until m=0, the prefix for next[j+1]=0
if s[0]=s[j], then next[j+1]=1 */

func KMPSearch(haystack string, needle string) int {
	next := getNextTable(needle)
	hLen, nLen := len(haystack), len(needle)
	kLen := hLen - nLen

	for i := 0; i <= kLen; {
		j, k := i, 0
		for j < hLen && k < nLen && haystack[j] == needle[k] {
			j++
			k++
		}
		if k == nLen {
			return i
		}
		i += k - next[k]
	}
	return -1
}

func getNextTable(s string) []int {
	n := len(s)
	if n == 2 {
		return []int{-1, 0}
	}

	next := make([]int, n)
	next[0] = -1
	// k = lenth of longest prefix
	k := 0
	for j := 1; j < n-1; j++ {
		for ; k > 0 && s[j] != s[k]; k = next[k] {
		}
		if s[j] == s[k] {
			k++
		}
		next[j+1] = k
	}
	return next
}

func search(haystack string, needle string) int {
	hLen, nLen := len(haystack), len(needle)
	kLen := hLen - nLen
	for i := 0; i <= kLen; i++ {
		j, k := i, 0
		for j < hLen && k < nLen && haystack[j] == needle[k] {
			j++
			k++
		}
		if k == nLen {
			return i
		}
	}
	return -1
}
