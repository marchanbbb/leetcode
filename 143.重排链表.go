/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func findMid(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	for fast!=nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		tmp := head.Next
		head.Next= prev
		prev= head
		head = tmp
	}
	return prev
}

func reorderList(head *ListNode)  {
	if head == nil {
		return
	}
	dummy1 := head
	mid := findMid(head)
	secondHalf := mid.Next
	l2 := reverseList(secondHalf)
	mid.Next = nil
	for dummy1 != nil {
		tmp := dummy1.Next
		dummy1.Next = l2
		l2 = tmp
		dummy1 = dummy1.Next
	}
}
// @lc code=end

