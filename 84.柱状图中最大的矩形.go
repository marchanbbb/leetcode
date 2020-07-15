/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	stack := make([]int, 0)
	max := 0
	for i := 0; i <= len(heights); i++ {
		cur := 0
		if i == len(heights) {
			cur = 0
		} else {
			cur = heights[i]
		}
		var h,w int
		for len(stack) != 0 && cur <= heights[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h = heights[pop]
			w= i
			if len(stack)!= 0 {
				peek := stack[len(stack)-1]
				w = i-peek-1
			}
		}
		max = getMax(max, h*w)
		stack = append(stack, i)
	}
	return max
}

func getMax(x, y int) int {
	if x < y {
		return y
	}
	return x
}
// @lc code=end

