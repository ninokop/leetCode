package main

import (
	"fmt"
)

func main() {
	origin := []int{2, 7, 11, 15}
	target := twoSum(origin, 9)
	fmt.Println(target)
}

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int, len(nums))
	for i, n := range nums {
		if id, ok := dict[n]; ok {
			return []int{id, i}
		}
		dict[target-n] = i
	}
	return nil
}
