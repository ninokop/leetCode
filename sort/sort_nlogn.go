package sort

func heapSort(d []int) {
	n := len(d)
	for i := n/2 - 1; i >= 0; i-- {
		down(d, i, n)
	}

	for i := n - 1; i > 0; i-- {
		swap(d, 0, i)
		down(d, 0, i)
	}
}

func down(d []int, lo, hi int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}

		if child+1 < hi && d[child+1] > d[child] {
			child++
		}

		if d[root] >= d[child] {
			return
		}

		swap(d, root, child)
		root = child
	}
}

func quickSort(d []int) {
	qsort(d, 0, len(d)-1)
}

func qsort(d []int, lo, hi int) {
	if lo < hi {
		pivot := doPivot(d, lo, hi)
		qsort(d, lo, pivot-1)
		qsort(d, pivot+1, hi)
	}
}

func doPivot(d []int, lo, hi int) int {
	pivot := d[lo]
	b, c := lo, hi

	for b < c {
		for b < c && d[c] >= pivot {
			c--
		}

		if b < c {
			d[b] = d[c]
			b++
		}

		for b < c && d[b] < pivot {
			b++
		}
		if b < c {
			d[c] = d[b]
			c--
		}
	}
	d[b] = pivot
	return b
}

func mergeSort(d []int) {
	sq := make([]int, len(d))
	msort(d, 0, len(d)-1, sq)
}

func msort(d []int, lo, hi int, sq []int) {
	if lo < hi {
		mid := lo + (hi-lo)/2
		msort(d, lo, mid, sq)
		msort(d, mid+1, hi, sq)
		merge(d, lo, mid, hi, sq)
	}
}

func merge(d []int, lo, mid, hi int, sq []int) {
	i, j, k := lo, mid+1, 0
	for ; i <= mid && j <= hi; k++ {
		if d[i] < d[j] {
			sq[k] = d[i]
			i++
		} else {
			sq[k] = d[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		sq[k] = d[i]
		k++
	}

	for ; j <= hi; j++ {
		sq[k] = d[j]
		k++
	}

	for i := 0; i < k; i++ {
		d[lo+i] = sq[i]
	}
}
