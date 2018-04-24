package sort

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	d := new(ListNode)
	d.Next = head
	for pre, cur := head, head.Next; cur != nil; {
		p, q := d, d.Next

		for q != cur && q.Val <= cur.Val {
			p, q = q, q.Next
		}
		if q != cur {
			pre.Next = cur.Next
			cur.Next = p.Next
			p.Next = cur
			cur = pre.Next
		} else {
			pre = pre.Next
			cur = pre.Next
		}
	}
	return d.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = slow.Next
	slow.Next = nil

	l1 := sortList(head)
	l2 := sortList(fast)

	return mergeTwoLists(l1, l2)
}
