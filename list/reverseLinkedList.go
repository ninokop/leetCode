package list

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		prev *ListNode
		next *ListNode
		cur  = head
	)

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	h1 := &ListNode{
		Val:  -1,
		Next: head,
	}

	prev := h1

	for i := 1; i < m; i++ {
		prev = prev.Next
	}
	h2 := prev
	prev = h2.Next
	cur := prev.Next

	for i := m; i < n; i++ {
		prev.Next = cur.Next
		cur.Next = h2.Next
		h2.Next = cur
		cur = prev.Next
	}
	return h1.Next
}
