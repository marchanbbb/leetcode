/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func findMiddle(head *ListNode) *ListNode {
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	} 
	return slow
}

func mergeTwoList(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for left != nil && right != nil {
		if left.Val > right.Val {
			tail.Next = right
			right = right.Next
		} else {
			tail.Next = left 
			left = left.Next
		}
		tail = tail.Next
	}
	if left != nil {
		tail.Next = left
	}
	if right != nil {
		tail.Next = right
	}
	return dummy.Next
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMiddle(head)
	l := mid.Next
	mid.Next = nil
	left := mergeSort(head)
	right := mergeSort(l)
	return mergeTwoList(left, right)
}

func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}
// @lc code=end

