/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(num int) []int {
	res := make([]int, num+1)
	for i:=1; i<=num; i++{
		res[i]= res[i/2] + (i%2)
	}
	return res
}
// @lc code=end

