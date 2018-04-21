package tree

func levelOrder(root *TreeNode) [][]int {
	out := make([][]int, 0)
	p := root
	q1 := newQueue()
	if p != nil {
		q1.Push(p)
	}

	for !q1.Empty() {

		n := q1.Len()
		level := make([]int, 0)
		for i := 0; i < n; i++ {
			cur := q1.Peek()
			q1.Pop()

			level = append(level, cur.Val)
			if cur.Left != nil {
				q1.Push(cur.Left)
			}
			if cur.Right != nil {
				q1.Push(cur.Right)
			}
		}

		out = append(out, level)
	}
	return out
}
