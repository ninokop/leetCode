package sort

/*
5 4 3 2 1
4 5 3 2 1
3 4 5 2 1
2 3 4 5 1
1 2 3 4 5
*/
func insertSort(data []int) {
	n := len(data)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && data[j] < data[j-1]; j-- {
			swap(data, j, j-1)
		}
	}
}

/*
5 4 3 2 1
4 3 2 1 5
3 2 1 4 5
2 1 3 4 5
1 2 3 4 5
*/
func bubbleSort(data []int) {
	flag := len(data)
	for flag > 0 {
		k := flag
		flag = 0
		for j := 1; j < k; j++ {
			if data[j] < data[j-1] {
				swap(data, j, j-1)
				flag = j
			}
		}

	}
}

/*
5 4 3 2 1
1 4 3 2 5
1 2 3 4 5
*/
func selectSort(data []int) {
	for i := 0; i < len(data); i++ {
		minIndex := i

		for j := i + 1; j < len(data); j++ {
			if data[j] < data[minIndex] {
				minIndex = j
			}
		}
		swap(data, i, minIndex)
	}
}

func swap(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}
