package list

import (
	"container/heap"
	gosort "sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	out := new(ListNode)
	p := out
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	} else {
		p.Next = l2
	}
	return out.Next
}

func mergeKListsBySC(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	p := lists[0]
	for i := 1; i < len(lists); i++ {
		p = mergeTwoLists(p, lists[i])
	}
	return p
}

func mergeKListsByRC(lists []*ListNode) *ListNode {
	n := len(lists)
	switch n {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		return mergeTwoLists(lists[0], lists[1])
	}

	B := mergeKListsByRC(lists[n/2:])
	A := mergeKListsByRC(lists[0 : n/2])

	return mergeTwoLists(A, B)
}

// MergeKList By Array
func mergeKListsByArray(lists []*ListNode) *ListNode {
	array := make([]int, 0)
	for _, list := range lists {
		array = append(array, listToArray(list)...)
	}
	gosort.Ints(array)
	return arrayToList(array)
}

func arrayToList(arr []int) *ListNode {
	head := new(ListNode)
	node := head

	n := len(arr)
	for i, v := range arr {
		node.Next = new(ListNode)
		node.Next.Val = v

		if i != n-1 {
			node = node.Next
		}
	}
	return head.Next
}

func listToArray(list *ListNode) []int {
	out := make([]int, 0)
	cur := list
	for cur != nil {
		out = append(out, cur.Val)
		cur = cur.Next
	}
	return out
}

// MergeKList By Heap
func mergeKListsByHeap(lists []*ListNode) *ListNode {
	q := new(Queue)
	heap.Init(q)

	for i := range lists {
		if lists[i] != nil {
			heap.Push(q, lists[i])
		}
	}

	if q.Len() == 0 {
		return nil
	}

	head := new(ListNode)
	cur := head
	for q.Len() > 0 {
		new := heap.Pop(q).(*ListNode)
		if new.Next != nil {
			heap.Push(q, new.Next)
		}

		cur.Next = new
		cur = cur.Next
	}
	return head.Next

}

type Queue []*ListNode

func (q Queue) Len() int           { return len(q) }
func (q Queue) Less(i, j int) bool { return q[i].Val < q[j].Val }
func (q Queue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *Queue) Push(x interface{}) { *q = append(*q, x.(*ListNode)) }
func (q *Queue) Pop() interface{} {
	x := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return x
}
