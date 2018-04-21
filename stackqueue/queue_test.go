package stackqueue

import "testing"
import "log"

func TestQueue(t *testing.T) {
	q := ConstructorQueue()
	q.Push(1)
	q.Push(2)
	log.Println(q.Pop())
	log.Println(q.Peek())
	log.Println(q.Empty())
}
