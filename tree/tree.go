package tree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack struct {
	queue []*TreeNode
	len   int
}

func newStack() stack {
	return stack{
		queue: make([]*TreeNode, 0),
		len:   0}
}

func (this *stack) Push(x *TreeNode) {
	this.queue = append(this.queue, x)
	this.len++

}

func (this *stack) Pop() *TreeNode {
	pop := this.queue[this.len-1]
	this.len--
	this.queue = this.queue[:this.len]
	return pop
}

func (this *stack) Top() *TreeNode { return this.queue[this.len-1] }
func (this *stack) Empty() bool    { return this.len == 0 }

type MyQueue struct {
	queue []*TreeNode
	len   int
}

func newQueue() *MyQueue {
	return &MyQueue{
		queue: make([]*TreeNode, 0),
		len:   0}
}

func (this *MyQueue) Push(x *TreeNode) {
	this.queue = append(this.queue, x)
	this.len++
}
func (this *MyQueue) Pop() *TreeNode {
	pop := this.queue[0]
	this.queue = this.queue[1:this.len]
	this.len--
	return pop
}
func (this *MyQueue) Peek() *TreeNode { return this.queue[0] }
func (this *MyQueue) Empty() bool     { return this.len == 0 }
func (this *MyQueue) Len() int        { return this.len }
