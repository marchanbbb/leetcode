/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i:=1; i<=len(s); i++{
		dp[i] = check(s[:i], wordDict)
		if dp[i] == true {
			continue
		}
		for k:=0; k<i; k++{
			if check(s[k:i], wordDict) && dp[k] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}
func check(s string, wordDict []string) bool {
	for _, s1 := range wordDict {
		if s == s1 {
			return true
		} 
	}
	return false
}
// @lc code=end

