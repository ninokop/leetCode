package tree

func preorderTraversalByRecursion(root *TreeNode) []int {
	nodes := make([]int, 0)
	if root == nil {
		return nodes
	}

	nodes = append(nodes, root.Val)
	nodes = append(nodes, preorderTraversalByRecursion(root.Left)...)
	return append(nodes, preorderTraversalByRecursion(root.Right)...)
}

func preorderTraversalByStack(root *TreeNode) []int {
	s := newStack()
	if root != nil {
		s.Push(root)
	}

	out := make([]int, 0)
	for !s.Empty() {
		p := s.Top()
		s.Pop()
		out = append(out, p.Val)

		if p.Right != nil {
			s.Push(p.Right)
		}
		if p.Left != nil {
			s.Push(p.Left)
		}
	}
	return out
}

// func preorderTraversal(root *TreeNode) []int {
// 	nodes := make([]int, 0)

// 	cur := root
// 	for cur != nil {
// 		if cur.Left != nil {
// 			node := cur.Left
// 			for node.Right != nil && node.Left != nil {
// 				node = node.Right
// 			}
// 		} else {

// 		}
// 	}
// }
