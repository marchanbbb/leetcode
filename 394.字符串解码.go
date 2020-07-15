/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	stack := make([]byte, 0)
	for i:=0; i< len(s);i++ {
		switch s[i] {
		case ']':
			tmp := make([]byte, 0)
			for len(stack) != 0&& stack[len(stack)-1]!= '[' {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				tmp=append(tmp, v)
			}
			stack = stack[:len(stack)-1]
			idx := 1
			for len(stack) >= idx&&stack[len(stack)-idx]>= '0' &&stack[len(stack)-idx]<='9' {
				idx++
			}
			num := stack[len(stack)-idx+1:]
			stack = stack[:len(stack)-idx+1]
			count, _ := strconv.Atoi(string(num))
			for j := 0; j<count; j++ {
				for l := len(tmp)-1;l>=0;l--{
					stack = append(stack, tmp[l])
				}
			}
		default:
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
// @lc code=end

