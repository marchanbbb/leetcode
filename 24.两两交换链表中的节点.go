/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	return helper(head)
}

func helper(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nextHead := head.Next.Next
	next := head.Next
	head.Next.Next = head
	head.Next = helper(nextHead)
	return next
}
// @lc code=end

