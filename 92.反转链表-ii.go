/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{}
	dummy.Next = head
	head = dummy
	i := 1
	for i<m {
		head=head.Next
		i++
	}
	prevNode := head
	head = head.Next
	mid := head
	var next *ListNode
	for i<=n && head != nil {
		tmp := head.Next
		head.Next = next
		next = head
		head = tmp
		i++
	}
	prevNode.Next = next
	mid.Next = head
	return dummy.Next
}
// @lc code=end

