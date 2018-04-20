package main

import (
	"container/heap"
	"log"
)

func main() {
	cases := [][]int{
		{2, 3, 5, 43, 21, 10, 33},
		{3, 2, 1, 5, 6, 4},
	}

	for _, s := range cases {
		log.Println(findKthLargest(s, 5))
	}
}

// //Kbig use partition
func findKthLargest(data []int, k int) int {
	return qSelect(data, 0, len(data)-1, k)
}

func qSelect(data []int, lo, hi, k int) int {
	if lo < hi {
		mid := partition(data, lo, hi)

		pLen := mid - lo + 1
		if pLen > k {
			return qSelect(data, lo, mid-1, k)
		} else if pLen < k {
			return qSelect(data, mid+1, hi, k-pLen)
		} else {
			return data[mid]
		}
	}
	return data[lo]
}

func partition(s []int, lo, hi int) int {
	i, j, pivot := lo, hi, s[lo]
	for i < j {
		for ; s[j] < pivot && i < j; j-- {
		}
		if i < j {
			s[i] = s[j]
			i++
		}
		for ; s[i] >= pivot && i < j; i++ {
		}
		if i < j {
			s[j] = s[i]
		}
	}
	s[i] = pivot
	return i
}

// Kbig use heap sort with my first try
func findKthLargestUseHeapSort(nums []int, k int) int {
	n := len(nums)

	for i := n/2 - 1; i >= 0; i-- {
		siftDown(nums, i, n)
	}

	for i := 0; i < k-1; i++ {
		nums[0], nums[n-1-i] = nums[n-1-i], nums[0]
		siftDown(nums, 0, n-1-i)
	}
	return nums[0]
}

func siftDown(nums []int, lo, hi int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}

		if child+1 < hi && nums[child+1] > nums[child] {
			child++
		}

		if nums[root] >= nums[child] {
			return
		}
		nums[root], nums[child] = nums[child], nums[root]
		root = child
	}
}

//Kbig use container heap package
func findKthLargestByContainerHeap(data []int, k int) int {
	// assume k is valid.
	h := make(IntHeap, k)
	f := true
	for i, v := range data {
		if i < k {
			h[i] = v
		} else if f {
			heap.Init(&h)
			f = false
		}
		if !f && h[0] < v {
			h[0] = v
			heap.Fix(&h, 0)
		}
	}
	if f {
		heap.Init(&h)
	}
	return h[0]
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[j], h[i] = h[i], h[j]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}
