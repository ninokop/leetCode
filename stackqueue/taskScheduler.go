package stackqueue

import (
// "sort"
)

func leastInterval(tasks []byte, n int) int {
	cnt := make([]int, 26)
	mx, mxCnt := 0, 0 // mx means A appears most frequently, mxCnt means how many alpha has mx cnt
	for _, task := range tasks {
		cnt[task-'A']++

		if mx == cnt[task-'A'] {
			mxCnt++
		} else if mx < cnt[task-'A'] {
			mx = cnt[task-'A']
			mxCnt = 1
		}
	}

	partCnt := mx - 1
	partLen := n - (mxCnt - 1)
	emptySlots := partCnt * partLen

	taskLeft := len(tasks) - mx*mxCnt
	idles := max(0, emptySlots-taskLeft)
	return len(tasks) + idles
}
