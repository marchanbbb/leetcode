/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
  numsM := map[int]bool{}
  for _, i := range nums {
	  if numsM[i] == true {
		  return true
	  }
	  numsM[i] = true
  }
  return false
}
// @lc code=end

