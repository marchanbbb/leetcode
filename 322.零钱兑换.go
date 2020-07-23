/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i:=1;i<=amount;i++{
        dp[i]=amount+1
    }
	for i:=1; i<=amount; i++ {
		for j:=0; j<len(coins); j++ {
			if i - coins[j] >= 0{
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}

func min(a, b int) int{
	if a > b {
		return b
	}
	return a
}
// @lc code=end

