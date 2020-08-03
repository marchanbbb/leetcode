/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	n2m := []string{"abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"}
	result := []string{""}
	for i:=0; i<len(digits); i++ {
		resultL := len(result)
		for j := 0; j<resultL; j++ {
			s := result[j]
			for k, c := range n2m[digits[i]-'2'] {
				v := s + string(c)
				if k == 0 {
					result[j] = v
				} else {
					result = append(result, v)
				}
			}
		}
	}
	return result
}
// @lc code=end

