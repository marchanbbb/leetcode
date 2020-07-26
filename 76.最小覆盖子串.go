/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	m := make(map[byte]int)
	matchM := make(map[byte]int)
	for i := 0; i<len(t); i++ {
		m[t[i]]++
	}
	var l, r, start, end, match int
	min := math.MaxInt64
	var c byte
	for r < len(s) {
		c = s[r]
		if m[c]!=0{
			matchM[c]++
			if matchM[c] == m[c] {
				match++
			}
		}
		r++
		for match == len(m) {
			if min > r-l {
				min = r-l
				start = l
				end = r
			}
			c = s[l]
			if m[c] != 0 {
				if m[c] == matchM[c] {
					match--
				}
				matchM[c]--
			}
			l++
		}
	}
	if min == math.MaxInt64{
		return ""
	}
	return s[start:end]
}
// @lc code=end

