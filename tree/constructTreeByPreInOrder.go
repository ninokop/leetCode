package tree

func buildTreeByPreInOrder(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	inRootPos := find(inorder, preorder[0])

	root.Left = buildTree(preorder[1:inRootPos+1], inorder[0:inRootPos])
	root.Right = buildTree(preorder[inRootPos+1:], inorder[inRootPos+1:])
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
