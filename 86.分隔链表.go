/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{}
	tmpList := &ListNode{}
	tmpListTail := tmpList
	dummy.Next = head
	head = dummy
	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
		} else {
			tmpListTail.Next = head.Next
			head.Next = head.Next.Next
			tmpListTail = tmpListTail.Next
		}
	}
	tmpListTail.Next = nil
	head.Next = tmpList.Next
	return dummy.Next
}
// @lc code=end

