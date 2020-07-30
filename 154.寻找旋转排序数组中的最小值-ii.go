/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start
func findMin(nums []int) int {
	if len(nums) == 0 {
        return -1
    }
	var start, end, mid int
	end = len(nums)-1
	for start+1 < end {
		for start < end && nums[end-1] == nums[end] {
			end--
		}
		for start < end && nums[start] == nums[start+1] {
			start++
		}
		mid = start + (end-start)/2
		if nums[mid] <= nums[end] {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] > nums[end] {
        return nums[end]
    }
    return nums[start]
}
// @lc code=end

