/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 0 {
        return -1
    }
	start, end := 0, len(nums)-1
	for start + 1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target{
			return mid
		}
		if nums[start] < nums[mid] {
			if nums[start] <= target && target <= nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else if nums[end] > nums[mid] {
			if nums[end] >= target && nums[mid] <= target {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if nums[start] == target {
        return start
    } else if nums[end] == target {
        return end
    }
	return -1
}
// @lc code=end

