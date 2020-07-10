/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow := head
	fast := head.Next
	for fast!= nil && fast.Next !=nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	secondHalf := slow.Next
	slow.Next = nil
	reverser := reverse(secondHalf)
	for head != nil && reverser != nil {
		if head.Val != reverser.Val {
			return false
		}
		head = head.Next
		reverser = reverser.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prevNode *ListNode
	for head !=nil {
		tmp := head.Next
		head.Next = prevNode
		prevNode = head
		head = tmp
	}
	return prevNode
}
// @lc code=end

