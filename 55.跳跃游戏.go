/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	max := 0
	for i, jump := range nums{
		if max >= i && max < i+jump {
			max = i+jump
		}
	}
	return len(nums)-1 <= max
}
/* dp解法
func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n+1)
	dp[0] = true
	for i:=1; i<n; i++{
		for j:=0; j<i; j++{
			if dp[i-1] && nums[j]+j >= i {
				dp[i] = true
			}
		}
	}
	return dp[n-1]
}
*/
// @lc code=end

