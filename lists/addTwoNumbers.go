package list

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	out := new(ListNode)
	carry := 0
	cur := out
	for l1 != nil || l2 != nil {
		cur.Next = new(ListNode)
		cur = cur.Next

		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		cur.Val = sum % 10
		carry = sum / 10

	}

	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return out.Next
}
