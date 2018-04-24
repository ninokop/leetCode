package sort

import "testing"
import "log"

func TestInsertSort(t *testing.T) {
	cases := [][]int{
		{4, 3, 6, 7, 8, 2, 3},
		{9, 23, 5, 6, 7, 12, 3, 12},
		{2},
		{3, 1},
	}

	for i := range cases {
		insertSort(cases[i])
		log.Println(cases[i])
	}
}

func TestBubbleSort(t *testing.T) {
	cases := [][]int{
		{4, 3, 6, 7, 8, 2, 3},
		{9, 23, 5, 6, 7, 12, 3, 12},
		{2},
		{3, 1},
	}

	for i := range cases {
		bubbleSort(cases[i])
		log.Println(cases[i])
	}
}

func TestSelectSort(t *testing.T) {
	cases := [][]int{
		{4, 3, 6, 7, 8, 2, 3},
		{9, 23, 5, 6, 7, 12, 3, 12},
		{2},
		{3, 1},
	}

	for i := range cases {
		selectSort(cases[i])
		log.Println(cases[i])
	}
}

func TestHeapSort(t *testing.T) {
	cases := [][]int{
		{4, 3, 6, 7, 8, 2, 3},
		{9, 23, 5, 6, 7, 12, 3, 12},
		{2},
		{3, 1},
	}

	for i := range cases {
		heapSort(cases[i])
		log.Println(cases[i])
	}
}

func TestQSort(t *testing.T) {
	cases := [][]int{
		{4, 3, 6, 7, 8, 2, 3},
		{9, 23, 5, 6, 7, 12, 3, 12},
		{2},
		{3, 1},
	}

	for i := range cases {
		quickSort(cases[i])
		log.Println(cases[i])
	}
}

func TestMergeSort(t *testing.T) {
	cases := [][]int{
		{4, 3, 6, 7, 8, 2, 3},
		{9, 23, 5, 6, 7, 12, 3, 12},
		{2},
		{3, 1},
	}

	for i := range cases {
		mergeSort(cases[i])
		log.Println(cases[i])
	}
}
