/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 1
	}
	for i := 0; i < l; i++ {
		if nums[i] <= 0 {
			nums[i] = l + 1
		}
	}
	for i := 0; i < l; i++ {
		j := nums[i]
		if j < 0 {
			j = j * -1
		}
		if j > 0 && j <= l {
			if nums[j-1] > 0 {
				nums[j-1] = -nums[j-1]
			}
		}
	}
	for i, n := range nums {
		if n > 0 {
			return i + 1
		}
	}
	return l + 1
}

// @lc code=end

