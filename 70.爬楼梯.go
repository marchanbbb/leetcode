/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	j, k, l := 0, 0, 1
	for i:=0;i<n;i++{
		j = k
		k = l
		l = j + k
	}
	return l
}
// @lc code=end

