/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	return calHeight(root).isBalanced
}

type ansType struct {
	height int
	isBalanced bool
}

func calHeight(root *TreeNode) (ans ansType) {
	if root == nil {
		ans.height = 0
		ans.isBalanced = true
		return
	}
	left := calHeight(root.Left)
	right := calHeight(root.Right)
	if left.isBalanced && right.isBalanced && abs(left.height, right.height) <= 1 {
		ans.isBalanced = true
	}
	ans.height = max(left.height, right.height)+1
	return
}
func abs(a, b int) int {
	if a>b {
		return a-b
	}
	return b-a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

