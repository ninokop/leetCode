package stack

func longestValidParentheses(s string) int {
	last, len := -1, 0
	st := newSimpleStack()
	for i := range s {
		if s[i] == '(' {
			st.Push(i)
		} else {
			if st.Empty() {
				last = i
			} else {
				st.Pop()
				if st.Empty() {
					len = max(len, i-last)
				} else {
					len = max(len, i-st.Top().(int))
				}
			}
		}
	}
	return len
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
