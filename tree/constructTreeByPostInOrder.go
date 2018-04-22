package tree

func buildTreeByPostInOrder(inorder []int, postorder []int) *TreeNode {
	inLen, postLen := len(inorder), len(postorder)
	if inLen == 0 || postLen == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[postLen-1]}
	inRootPos := find(inorder, postorder[postLen-1])

	root.Left = buildTree(inorder[0:inRootPos], postorder[0:inRootPos])
	root.Right = buildTree(inorder[inRootPos+1:], postorder[inRootPos:postLen-1])
	return root
}

func find(inorder []int, val int) int {
	for i := range inorder {
		if val == inorder[i] {
			return i
		}
	}
	return -1
}
