/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 1; i < m; i++ {
        grid[i][0] = grid[i][0] + grid[i-1][0]
    }
    for j := 1; j < n; j++ {
        grid[0][j] = grid[0][j] + grid[0][j-1]
    }
	for i:=1; i< m; i++{
		for j:=1; j<n; j++{
			grid[i][j] += min(grid[i][j-1], grid[i-1][j])
		}
	}
	return grid[m-1][n-1]
}
func min(a, b int) int{
	if a<b{
		return a
	}
	return b
}
// @lc code=end

