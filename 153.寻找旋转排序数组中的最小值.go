/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	var start, end, mid int 
	start, end = 0, len(nums)-1
	target := nums[end]
	for start+1 < end {
		mid = start + (end-start)/2
		if nums[mid] < target {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] < nums[end] {
		return nums[start]
	}
	return nums[end]
}
// @lc code=end

