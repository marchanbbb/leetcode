/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */

// @lc code=start
/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    start := 0
    end := n
    for start+1 < end {
        mid := start + (end - start)/2
        if isBadVersion(mid) {
            end = mid
        } else if isBadVersion(mid) == false {
            start = mid
        }
    }
    if isBadVersion(start) {
        return start
    }
    return end
}
// @lc code=end

