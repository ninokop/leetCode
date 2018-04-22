package stack

import (
	"strings"
)

func isValid(s string) bool {
	left, right := "([{", ")]}"

	st := newSimpleStack()
	for i := range s {
		if strings.LastIndexByte(left, s[i]) != -1 {
			st.Push(s[i])
		} else {
			if st.Empty() || st.Top().(byte) != left[strings.LastIndexByte(right, s[i])] {
				return false
			}
			st.Pop()
		}
	}
	return st.Empty()
}
