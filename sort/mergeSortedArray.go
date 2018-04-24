package sort

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
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
