/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	res := 1
	for i:=3; i<n; i++{
		if i & 1 == 0 {
			continue
		}
		sign := true
		for j:=3; j*j<=i; j+=2 {
			if i%j == 0 {
				sign = false
				break
			}
		}
		if sign {
			res++
		}
	}
	return res
}
// @lc code=end

