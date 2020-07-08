/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	l := midsearch(root)
	for i := len(l)-1; i > 0; i-- {
		if l[i] <= l[i-1] {
			return false
		}
	}
	return true
}

func midsearch(root *TreeNode) []int {
	var l []int
	if root == nil {
		return l
	}
	if root.Left != nil {
		l = append(midsearch(root.Left), l...)
	}
	l = append(l, root.Val)
	if root.Right != nil {
		l = append(l, midsearch(root.Right)...)
	}
	return l
}
// @lc code=end

