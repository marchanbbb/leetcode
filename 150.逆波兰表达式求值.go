/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	if len(tokens)==0{
		return 0
	}
	stack := make([]int, 0)
	for i:=0;i<len(tokens);i++ {
		switch tokens[i] {
		case "+","-","*","/":
			if len(stack) < 2 {
				return -1
			}
			a:= stack[len(stack)-2]
			b:= stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			result := 0
			switch tokens[i] {
			case "+":
				result = a+b
			case "-":
				result = a-b
			case "*":
				result = a*b
			case "/":
				result = a/b
			}
			stack = append(stack, result)
		default:
			val,_ := strconv.Atoi(tokens[i])
            stack = append(stack,val)
		}
	}
	return stack[0]
}
// @lc code=end

