/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	front []int
	back []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	if idx := len(this.back)-1; idx >= 0 {
		this.front = append(this.front, this.back[idx])
		this.back = this.back[:idx]
	}
	this.front = append(this.front, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	for idx := len(this.front)-1; idx>=0 ; idx-- {
		this.back = append(this.back, this.front[idx])
	}
	this.front = this.front[:0]
	res := this.back[len(this.back)-1]
	this.back = this.back[:len(this.back)-1]
	return res
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	for idx := len(this.front)-1; idx>=0 ; idx-- {
		this.back = append(this.back, this.front[idx])
	}
	this.front = this.front[:0]
	if len(this.back) == 0 {
		return 0
	}
	return this.back[len(this.back)-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.back) == 0 && len(this.front) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

