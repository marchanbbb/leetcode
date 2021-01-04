/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	l := len(nums)
	numsM := make(map[int]int, l)
	for i := 0; i < l; i++ {
		n := nums[i]
		if numsM[n] > 0 {
			if j := i - numsM[n] + 1; j <= k {
				return true
			}
		}
		numsM[n] = i + 1
	}
	return false
}

// @lc code=end

