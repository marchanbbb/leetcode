/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
	var start, end int
	row, col := len(matrix), len(matrix[0])
	end = row * col - 1
	for start+1 < end {
		mid := start + (end-start)/2
		val := matrix[mid/col][mid%col]
		if val > target {
            end = mid
        } else if val < target {
            start = mid
        } else {
            return true
        }
	}
	if matrix[start/col][start%col] == target || matrix[end/col][end%col] == target{
        return true
    }
    return false
}
// @lc code=end

