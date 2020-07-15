/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for len(stack)>0 || root!=nil {
		for root!=nil {
			stack = append(stack, root)
			root = root.Left
		}
		val := stack[len(stack)-1]
		result = append(result, val.Val)
		stack = stack[:len(stack)-1]
		root = val.Right
	}
	return result
}
// @lc code=end

