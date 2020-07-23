/**
 * @param m: An integer m denotes the size of a backpack
 * @param A: Given n items with size A[i]
 * @return: The maximum size
 */
 func backPack (m int, A []int) int {
    // write your code here
    dp := make([][]bool, len(A)+1)
    for i:=0; i<len(A)+1; i++{
        dp[i] = make([]bool, m+1)
    }
    dp[0][0] = true
    for i:=1; i<=len(A); i++{
        for j:=0; j<=m; j++{
            dp[i][j]=dp[i-1][j]
            if j - A[i-1] >= 0 && dp[i-1][j-A[i-1]] {
                dp[i][j]=true
            }
        }
    }
    for i:=m;i>=0;i--{
        if dp[len(A)][i] {
            return i
        }
    }
    return 0
}
