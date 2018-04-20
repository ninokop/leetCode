package list

import (
	"fmt"
	"log"
	"testing"
)

func TestReorderList(t *testing.T) {
	log.Println("TestReorderList")
	l1 := NewListNode(1)
	for i := 2; i <= 10; i++ {
		l1.PushBack(i)
	}

	l2 := reorderList(l1)
	for i := 1; i <= 10; i++ {
		n := l2.PopFront()
		fmt.Printf("%d ", n)
	}
	fmt.Printf("\n")
}
