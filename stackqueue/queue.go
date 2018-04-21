package stackqueue

type QueueByStacks struct {
	s1 MyStack
	s2 MyStack
}

/** Initialize your data structure here. */
func NewQueueByStacks() QueueByStacks {
	return QueueByStacks{
		s1: NewStack(),
		s2: NewStack(),
	}
}

/** Push element x onto stack. */
func (this *QueueByStacks) Push(x int) {
	this.s1.Push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *QueueByStacks) Pop() int {
	if this.s2.Empty() {
		for !this.s1.Empty() {
			this.s2.Push(this.s1.Pop())
		}
	}
	return this.s2.Pop()
}

/** Get the top element. */
func (this *QueueByStacks) Top() int {
	if this.s2.Empty() {
		for !this.s1.Empty() {
			this.s2.Push(this.s1.Pop())
		}
	}
	return this.s2.Top()
}

/** Returns whether the stack is empty. */
func (this *QueueByStacks) Empty() bool {
	return this.s1.Empty() && this.s2.Empty()
}

type MyQueue struct {
	queue []int
	len   int
}

/** Initialize your data structure here. */
func ConstructorQueue() MyQueue {
	return MyQueue{
		queue: make([]int, 0),
		len:   0}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
	this.len++
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	pop := this.queue[0]
	this.queue = this.queue[1:this.len]
	this.len--
	return pop
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.queue[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.len == 0
}
