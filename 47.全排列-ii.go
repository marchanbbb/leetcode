/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	used := make([]bool, len(nums))
	sort.Ints(nums)
	backtrace(nums, used, list, &result)
	return result
}

func backtrace(nums []int, used []bool, list []int, result *[][]int) {
	if len(list) == len(nums) {
		ans := make([]int, len(nums))
		copy(ans, list)
		*result = append(*result, ans)
		return
	}
	for i := 0; i<len(nums); i++ {
		if used[i] == true {
			continue
		}
		if i != 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		list = append(list, nums[i])
		used[i] = true
		backtrace(nums, used, list, result)
		used[i] = false
		list = list[0:len(list)-1]
	}
}
// @lc code=end

