/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count:=0
	for i:= 0; i< len(grid);i++ {
		for j:=0; j<len(grid[0]);j++{
			if dfs(grid, i, j)>=1 {
				count++
			} 
		}
	}
	return count
}

func dfs(grid [][]byte, i,j int) int {
	if i<0 || j <0 || i >= len(grid) || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == '1' {
		grid[i][j] = 0
		return dfs(grid,i-1,j)+dfs(grid,i+1,j)+dfs(grid,i,j-1)+dfs(grid,i,j+1)+1
	}
	return 0
}
// @lc code=end

