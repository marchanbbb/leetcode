/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)
	list := make([]int, 0)
	backtrack(candidates, list, target, 0, &result)
	return result
}

func backtrack(candidates, list []int, target, pos int,result *[][]int) {
	if target == 0 {
		tmp := make([]int, len(list))
		copy(tmp, list)
		*result = append(*result, tmp)
	}
	for i := pos; i < len(candidates); i++ {
		v := candidates[i]
		if target < v {
			break
		}
		backtrack(candidates, append(list, v), target-v, i, result)
	}
}

// @lc code=end

