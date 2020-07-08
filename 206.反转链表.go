/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prevNode *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prevNode
		prevNode = head
		head = temp
	}
	return prevNode
}
// @lc code=end

