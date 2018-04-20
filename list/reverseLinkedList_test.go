package list

import "testing"

func TestReverseList(t *testing.T) {
	l1 := NewListNode(1)
	for i := 2; i <= 10; i++ {
		l1.PushBack(i)
	}

	l2 := reverseList(l1)
	for i := 10; i > 10; i-- {
		if n := l2.PopFront(); n != i {
			t.Errorf("l2.PopFront = %d, want %d", n, i)
		}
	}
}

func TestReverseBetween(t *testing.T) {
	l1 := NewListNode(1)
	for i := 2; i <= 10; i++ {
		l1.PushBack(i)
	}

	l2 := reverseBetween(l1, 6, 10)
	for i := 1; i <= 5; i++ {
		if n := l2.PopFront(); n != i {
			t.Errorf("l2.PopFront = %d, want %d", n, i)
		}
	}

	for i := 10; i > 5; i-- {
		if n := l2.PopFront(); n != i {
			t.Errorf("l2.PopFront = %d, want %d", n, i)
		}
	}

	l3 := NewListNode(5)
	l4 := reverseBetween(l3, 1, 1)
	if n := l4.PopFront(); n != 5 {
		t.Errorf("l4.PopFront = %d, want 5", n)
	}

	l5 := NewListNode(3)
	l5.PushBack(5)
	l6 := reverseBetween(l5, 1, 1)
	if n := l6.PopFront(); n != 3 {
		t.Errorf("l4.PopFront = %d, want 5", n)
	}

	if n := l6.PopFront(); n != 5 {
		t.Errorf("l4.PopFront = %d, want 5", n)
	}

}
