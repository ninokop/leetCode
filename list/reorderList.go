package list

func reorderList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	s, f, h1 := head, head, head

	for f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	h2 := reverseOf(s)
	s.Next = nil

	for h1.Next != nil {
		tmp := h1.Next
		h1.Next = h2
		h2 = h2.Next
		h1.Next.Next = tmp
		h1 = tmp
	}

	if h2 != nil {
		h1.Next = h2
	}
	return head
}

func reverseOf(prev *ListNode) *ListNode {
	cur := prev.Next
	next := cur.Next

	for next != nil {
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
		next = cur.Next
	}
	return prev.Next
}
