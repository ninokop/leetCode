package list

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	l1 := NewListNode(2)
	l1.PushBack(4)
	l1.PushBack(3)

	l2 := NewListNode(5)
	l2.PushBack(6)
	l2.PushBack(4)

	l3 := addTwoNumbers(l1, l2)
	if n := l3.PopFront(); n != 7 {
		t.Errorf("l3.PopFront = %d, want 7", n)
	}
	if n := l3.PopFront(); n != 0 {
		t.Errorf("l3.PopFront = %d, want 0", n)
	}
	if n := l3.PopFront(); n != 8 {
		t.Errorf("l3.PopFront = %d, want 8", n)
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	l1 := NewListNode(5)
	l2 := NewListNode(5)

	l3 := addTwoNumbers(l1, l2)
	if n := l3.PopFront(); n != 0 {
		t.Errorf("l3.PopFront = %d, want 0", n)
	}
	if n := l3.PopFront(); n != 1 {
		t.Errorf("l3.PopFront = %d, want 1", n)
	}
}
