/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy
	var rmVal int
	for head.Next!=nil && head.Next.Next!=nil {
		if head.Next.Val == head.Next.Next.Val {
			rmVal = head.Next.Val
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}
// @lc code=end

