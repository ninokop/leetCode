package list

import "testing"

func TestRemoveNthNode(t *testing.T) {
	l1 := NewListNode(1)
	for i := 2; i <= 10; i++ {
		l1.PushBack(i)
	}

	l2 := removeNthFromEnd(l1, 2)
	for i := 1; i <= 8; i++ {
		if n := l2.PopFront(); n != i {
			t.Errorf("l2.PopFront = %d, want %d", n, i)
		}
	}

	if n := l2.PopFront(); n != 10 {
		t.Errorf("l2.PopFront = %d, want 10", n)
	}
}
