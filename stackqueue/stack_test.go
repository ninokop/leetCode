package stackqueue

import "testing"
import "log"

func TestStack(t *testing.T) {
	stack := ConstructorStack()
	stack.Push(1)
	stack.Push(2)
	log.Println(stack.Pop())
	log.Println(stack.Top())
	log.Println(stack.Empty())
}

func TestStackByQueues(t *testing.T) {
	stack := NewStackByQueues()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	log.Println(stack.Pop())
	log.Println(stack.Top())
	stack.Push(4)
	stack.Push(5)
	log.Println(stack.Pop())
	log.Println(stack.Empty())
}
