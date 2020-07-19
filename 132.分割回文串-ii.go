/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */

// @lc code=start
func minCut(s string) int {
	dp := make([]int, len(s)+1)
	dp[0]=-1
	for i:=1; i<=len(s); i++ {
		dp[i]=i-1
		for j:=0; j<i; j++ {
			if isPal(s[j:i]) && dp[i] > dp[j]+1{
				dp[i]=dp[j]+1
			}
		}
	}
	return dp[len(s)]
}

func isPal(s string) bool {
	for i ,n := 0,len(s)-1; i<len(s)/2; i++ {
		if s[i] != s[n-i] {
			return false
		}
	}
	return true
}
// @lc code=end

