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

func swapPairs(head *ListNode) *ListNode {
	d := new(ListNode)
	d.Next = head

	for cur := d; cur.Next != nil && cur.Next.Next != nil; {
		n1, n2 := cur.Next, cur.Next.Next
		n1.Next, n2.Next, cur.Next = n2.Next, n1, n2
		cur = n1
	}
	return d.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	d := &ListNode{0, head}
	prev := d
	end := prev.Next
	for end != nil {
		for i := 1; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		prev = reverse(prev, end.Next) //prev[...]end
		end = prev.Next
	}
	return d.Next
}

func reverse(prev, end *ListNode) *ListNode {
	if prev.Next == end {
		return end
	}
	cur := prev.Next
	next := cur.Next

	for next != end {
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
		next = cur.Next
	}
	return cur
}
