给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
————————————————
版权声明：本文为CSDN博主「Allen_Xu11」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/XUCHEN1230/article/details/104865519

func maxSubArray(nums []int) int {
    //数组求子序列最值问题:
    //F(i) = max(F(j)) + nums[i]
    dp := make([]int,len(nums))
    dp[0] = nums[0]
    res := nums[0]
    for i:=1;i<len(nums);i++ {
        dp[i] = int(math.Max(float64(dp[i-1]+nums[i]),float64(nums[i])))
        if dp[i] > res {
            res = dp[i]
        }
    }
    return res
}