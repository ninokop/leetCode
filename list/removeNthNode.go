package list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	d := &ListNode{-1, head}
	h1 := d
	h2 := d

	for i := 0; i < n; i++ {
		h2 = h2.Next
	}

	for h2.Next != nil {
		h1 = h1.Next
		h2 = h2.Next
	}

	remove := h1.Next
	h1.Next = h1.Next.Next
	remove.Next = nil

	return d.Next
}
