package stackqueue

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	out := heights[0]
	s := NewStack()

	for i := 0; i < len(heights); i++ {
		if s.Empty() || heights[s.Top()] <= heights[i] {
			s.Push(i)
		} else {
			tmp := s.Top()
			s.Pop()

			if s.Empty() {
				out = max(out, heights[tmp]*i)
			} else {
				out = max(out, heights[tmp]*(i-s.Top()-1))
			}

			i--
		}
	}
	return out
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
