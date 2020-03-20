你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

示例 1:

输入: [1,2,3,1]
输出: 4
解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 2:

输入: [2,7,9,3,1]
输出: 12
解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
————————————————
版权声明：本文为CSDN博主「Allen_Xu11」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/XUCHEN1230/article/details/104795649

func rob(nums []int) int {
    
    if len(nums) == 0 {
        return 0
    } else if len(nums) == 1 {
        return nums[0]
    } else if len(nums) == 2 {
        return max(nums[0],nums[1])
    }
 
    dp := make([]int,len(nums))
 
    dp[0] = nums[0]
    dp[1] = nums[1]
 
    for i:=2;i<len(nums);i++ {
        tmp := 0
        for j:=0;j<i-1;j++ {
            tmp = max(dp[j]+nums[i],tmp)
        }
        dp[i] = max(tmp,dp[i-1])
    }
 
    return dp[len(nums)-1]
}
func max (a,b int) int {
    if a > b {
        return a
    } else {
        return b
    }
    return 0
}