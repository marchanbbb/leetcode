/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	start, end := 0, len(nums)-1
	for start+1<end {
		for start<end && nums[end-1] == nums[end] {
			end--
		}
		for start<end && nums[start] == nums[start+1] {
			start++
		}
		mid := start+ (end-start)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] >= nums[start] {
			if target >= nums[start] && target <= nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else {
			if target >= nums[mid] && target <= nums[end] {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if nums[end] == target || nums[start] == target {
		return true
	}
	return false
}
// @lc code=end

