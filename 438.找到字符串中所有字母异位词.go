/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	matchC := make(map[byte]int)
	needC := make(map[byte]int)
	ans := make([]int, 0)
	for i:=0; i<len(p); i++ {
		needC[p[i]]++
	}
	var l, r, match int
	for r < len(s) {
		c := s[r]
		r++
		if needC[c] != 0 {
			matchC[c]++
			if matchC[c]==needC[c] {
				match++
			}
		}
		for r-l>=len(p) {
			if r-l == len(p) && match==len(needC) {
				ans = append(ans, l)
			}
			c1 := s[l]
			l++
			if needC[c1] != 0 {
				if matchC[c1] == needC[c1] {
					match--
				}
				matchC[c1]--
			}
		}
	}
	return ans
}
// @lc code=end

