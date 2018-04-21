package main

import (
	"log"
	"math"
	"strings"
)

func main() {

	str := []string{
		"  -43",
		"  010",
		" +001+2a43",
		"2147483648",
		"-2147483648",
		" -321",
		"   -119198374x",
		"92342453453456456",
	}
	for _, i := range str {
		a := myAtoi(i)
		log.Printf("%d", a)
	}
}

func myAtoi(str string) int {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return 0
	}

	flag := 1
	if str[0] == '-' {
		flag = -1
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}

	var result int = 0
	for i := 0; i < len(str); i++ {
		var c byte = str[i]
		if '0' <= c && c <= '9' {
			result = result*10 + int(c-'0')
			switch {
			case flag == 1 && result > math.MaxInt32:
				return math.MaxInt32
			case flag == -1 && -result < math.MinInt32:
				return math.MinInt32
			}
		} else {
			break
		}

	}

	return int(flag * result)
}
