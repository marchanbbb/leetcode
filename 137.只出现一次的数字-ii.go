/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
func singleNumber(nums []int) int {
	result := 0
	for i:=0; i<64; i++{
		sum := 0
		for j:=0;j<len(nums);j++{
			sum += (nums[j]>>i) & 1
		}
		result |= (sum%3) << i
	}
	return result
}
// @lc code=end

