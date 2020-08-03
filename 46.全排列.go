/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	used := make([]bool, len(nums))
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
		list = append(list, nums[i])
		used[i] = true
		backtrace(nums, used, list, result)
		used[i] = false
		list = list[0:len(list)-1]
	}
}
// @lc code=end

