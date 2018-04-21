package list

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func detectCycle(head *ListNode) *ListNode {
	s1, f := head, head
	for f != nil && f.Next != nil {
		s1 = s1.Next
		f = f.Next.Next
		if s1 == f {
			s2 := head

			for s2 != s1 {
				s2 = s2.Next
				s1 = s1.Next
			}
			return s2
		}
	}
	return nil
}
