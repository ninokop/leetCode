package stackqueue

type StackByQueues struct {
	q1 MyQueue
	q2 MyQueue
}

/** Initialize your data structure here. */
func NewStackByQueues() StackByQueues {
	return StackByQueues{
		q1: ConstructorQueue(),
		q2: ConstructorQueue(),
	}
}

/** Push element x onto stack. */
func (this *StackByQueues) Push(x int) {
	if this.q1.Empty() {
		this.q2.Push(x)
		return
	}
	this.q1.Push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *StackByQueues) Pop() int {
	if this.q1.Empty() {
		for !this.q2.Empty() {
			p := this.q2.Pop()
			if this.q2.Empty() {
				return p
			}
			this.q1.Push(p)
		}
	} else {
		for !this.q1.Empty() {
			p := this.q1.Pop()
			if this.q1.Empty() {
				return p
			}
			this.q2.Push(p)
		}
	}
	return -1
}

/** Get the top element. */
func (this *StackByQueues) Top() int {
	if this.q1.Empty() {
		for !this.q2.Empty() {
			p := this.q2.Pop()
			this.q1.Push(p)
			if this.q2.Empty() {
				return p
			}
		}
	} else {
		for !this.q1.Empty() {
			p := this.q1.Pop()
			this.q2.Push(p)
			if this.q1.Empty() {
				return p
			}
		}
	}
	return -1
}

/** Returns whether the stack is empty. */
func (this *StackByQueues) Empty() bool {
	return this.q1.Empty() && this.q2.Empty()
}

type MyStack struct {
	queue []int
	len   int
}

/** Initialize your data structure here. */
func NewStack() MyStack {
	return MyStack{
		queue: make([]int, 0),
		len:   0}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
	this.len++

}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	pop := this.queue[this.len-1]
	this.len--
	this.queue = this.queue[:this.len]
	return pop
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue[this.len-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.len == 0
}
