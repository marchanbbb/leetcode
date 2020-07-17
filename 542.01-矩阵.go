/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */

// @lc code=start
func updateMatrix(matrix [][]int) [][]int {
	m,n := len(matrix), len(matrix[0])
	dp := make([][]int,0)
	for i:=0;i<m;i++{
		dp =append(dp, make([]int, n))
	}
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if matrix[i][j] == 0{
				dp[i][j] = 0
			} else {
				dp[i][j] = 1000000
			}
		}
	}

	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if i-1>=0{
				dp[i][j]=GetMin(dp[i][j], dp[i-1][j]+1)
			}
			if j-1>=0{
				dp[i][j]=GetMin(dp[i][j], dp[i][j-1]+1)
			}
		}
	}
	for i:=m-1;i>=0;i--{
		for j:=n-1;j>=0;j--{
			if i+1<m{
				dp[i][j]=GetMin(dp[i][j], dp[i+1][j]+1)
			}
			if j+1<n{
				dp[i][j]=GetMin(dp[i][j], dp[i][j+1]+1)
			}
		}
	}
	return dp
}

func GetMin(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

