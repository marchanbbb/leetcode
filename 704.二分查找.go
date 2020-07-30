/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	var mid int
	for l<=r {
		mid = (l+r)/2
		if nums[mid] > target {
			r = mid-1
		} else if nums[mid] < target {
			l = mid+1
		} else {
			return mid
		}
	}
	return -1
}
// @lc code=end

