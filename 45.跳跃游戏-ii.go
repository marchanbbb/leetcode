/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	maxStep := 0
	maxPosition := 0
	end := 0
	for i := 0; i <len(nums)-1; i++{
		if maxPosition < nums[i]+ i{
			maxPosition = nums[i] + i
		}
		if end == i {
			end = maxPosition
			maxStep++
		}
	}
	return maxStep
}



/* dp
func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = 0
	for i:=1; i<n; i++{
		dp[i] = i
		for j:=0; j<i; j++{
			if nums[j] + j >= i {
				dp[i] = min(dp[j]+1, dp[i])
			}
		}
	}
	return dp[n-1]
}
func min(a, b int) int{
	if a > b {
		return b
	}
	return a
}
*/
// @lc code=end

