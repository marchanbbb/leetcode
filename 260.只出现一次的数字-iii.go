/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
func singleNumber(nums []int) []int {
	ab := 0
	for i:=0; i<len(nums); i++{
		ab ^= nums[i]
	}
	ab = (ab&(ab-1))^ab
	result := []int{0,0}
	for i:=0;i<len(nums);i++{
		if nums[i]&ab==0{
			result[0]^=nums[i]
		} else {
			result[1]^=nums[i]
		}
	}
	return result
}
// @lc code=end

