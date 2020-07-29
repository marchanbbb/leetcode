/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	var l, r, match int
	matchChar := make(map[byte]int)
	needChar := make(map[byte]int)
	for i:=0; i<len(s1); i++{
		needChar[s1[i]]++
	}
	for r<len(s2) {
		c := s2[r]
		r++
		if needChar[c] != 0 {
			matchChar[c]++
			if matchChar[c]==needChar[c]{
				match++
			}
		}
		for r-l >= len(s1) {
			if match == len(needChar) {
				return true
			}
			c = s2[l]
			l++
			if needChar[c] != 0 {
				if matchChar[c] == needChar[c] {
					match--
				}
				matchChar[c]--
			}
		}
	}
	return false
}
// @lc code=end

