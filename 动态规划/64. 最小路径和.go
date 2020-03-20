64. 最小路径和
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。


func minPathSum(grid [][]int) int {
    // F(x,y) = min(F(x-1,y),F(x,y-1)) + grid[x][y]

    m := len(grid)
    n := len(grid[0])

    dp := make([][]int,m)

    for k,_ := range grid {
        t := make([]int,n)
        dp[k] = t
    }

    dp[0][0] = grid[0][0]

    for i:=0;i<m;i++ {
        for j:=0;j<n;j++ {
            p := 0
            q := 0
            if i-1 >= 0 {
                p = dp[i-1][j] + grid[i][j]
            }
            if j-1 >= 0 {
                q = dp[i][j-1] + grid[i][j]
            }
            if p>0 && q>0 {
                dp[i][j] = Min(p,q)
            } else if p==0 && q>0 {
                dp[i][j] = q
            } else if p>0 && q==0 {
                dp[i][j] = p
            }
        }
    }

    return dp[m-1][n-1]
}


func Min(a,b int) int {
    if a<b {
        return a
    } else {
        return b
    }
}
