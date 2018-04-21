package main

import (
	"fmt"
	"math"
)

func main() {
	x := -123
	y := reverse2(x)
	fmt.Println(y)
}

func reverse1(x int) int {
	sum := 0
	for {
		if x == 0 {
			break
		}
		var t int32
		t = int32(sum*10 + x%10)
		if t/10 != int32(sum) {
			return 0
		}
		sum = int(t)
		x = x / 10
	}
	return sum
}

func reverse2(x int) int {
	var sum int32 = 0
	for {
		if x == 0 {
			break
		}
		if sum > math.MaxInt32/10 || sum < math.MinInt32/10 {
			return 0
		}
		sum = sum*10 + int32(x%10)
		x = x / 10
	}
	return int(sum)
}
