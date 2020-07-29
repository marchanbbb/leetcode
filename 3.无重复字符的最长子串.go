/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
		return 0
	}
	match := make(map[byte]int)
	l, r, res := 0, 0, 1
	for r<len(s) {
		match[s[r]]++
		for match[s[r]] > 1 {
			match[s[l]]--
			l++
		}
		r++
		res = max(res, r-l)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

