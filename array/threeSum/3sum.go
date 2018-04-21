package main

import (
	"log"
	"sort"
)

func main() {
	cases := [][]int{
		{-2, 0, 0, 2, 2},
		{-1, 0, 1, 2, -1, -4},
		{-2, 0, 1, 1, 2},
		{0, 0, 0},
	}

	for _, s := range cases {
		log.Println(threeSum(s))
	}
}

func threeSum(nums []int) [][]int {
	sort.IntSlice(nums).Sort()

	n := len(nums)
	out := make([][]int, 0)
	for i := 0; i < n-2; {
		sum := 0 - nums[i]
		for j, k := i+1, n-1; j < k; {

			curSum := nums[j] + nums[k]
			if curSum > sum {
				k--
			} else if curSum < sum {
				j++
			} else {
				out = append(out, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for (j < k) && (nums[j] == nums[j-1]) {
					j++
				}

				for (j < k) && (nums[k] == nums[k+1]) {
					k--
				}
			}
		}
		i++
		for i < n-2 && nums[i] == nums[i-1] {
			i++
		}
	}
	return out
}
