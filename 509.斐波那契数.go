/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(N int) int {
	if N == 0 || N == 1 {
		return N
	}
	x, y, z := 1,0,0
	for i :=0; i<N-1; i++{
		z = x + y
		y = x
		x = z
	}
	return z
}
// @lc code=end

