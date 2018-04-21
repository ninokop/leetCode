package main

import (
	"fmt"
)

func main() {

	// n2 := []int{2, 2, 5, 7, 8, 13, 44}
	// n1 := []int{1, 2, 3, 7, 9, 15, 0, 0, 0, 0, 0, 0, 0}
	// merge(n1, 6, n2, 7)

	n1 := []int{4, 5, 6, 0, 0, 0}
	n2 := []int{1, 2, 3}
	merge(n1, 3, n2, 3)
	fmt.Println(n1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	j, k := n-1, m+n-1
	for i := m - 1; i >= 0 && j >= 0; k-- {
		if nums1[i] <= nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
