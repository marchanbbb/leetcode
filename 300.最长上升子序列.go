/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1{
		return n
	}
	dp := make([]int, n)
	dp[0]=1
	for i:=1; i<n; i++{
		dp[i]=1
		for k:=0; k<i;k++{
			if nums[k] < nums[i] && nums[k] < nums[i] {
				dp[i] = max(dp[i], dp[k]+1)
			}
		}
	}
	max := 0
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}
// @lc code=end

