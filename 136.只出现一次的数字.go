/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
func singleNumber(nums []int) int {
    result := 0
    for _, v := range nums {
        result = result^v
    }
    return result
}
// @lc code=end

