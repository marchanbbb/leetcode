/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	dummy := head
	for head != nil {
		copy := &Node{
			Val:  head.Val,
			Next: head.Next,
		}
		head.Next = copy
		head = copy.Next
	}
	head = dummy
	for head !=nil {
		if head.Random != nil {
            head.Next.Random = head.Random.Next
		}
		head = head.Next.Next
	}
	head = dummy
	copyRandomListHead := head.Next
	for head!=nil && head.Next != nil {
		tmp := head.Next
		head.Next = head.Next.Next
		head = tmp
	}
	return copyRandomListHead
}
// @lc code=end

