给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"



func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	//创建dp数组记录子串情况
	//y轴为length:子串长度
	//x轴为end:记录以这个下标为最后一个字符的子串
	//dp含义:每个格子表示-->以end结尾的长度为length的子串是否为回文子串
	dp := make([][]bool, len(s))
	y := 0
	x := 0
	for i := 0; i < len(s); i++ {
		tmp := make([]bool, len(s))
		if i == 0 {
			for j := 0; j < len(s); j++ {
				//初始化dp表
				tmp[j] = true
			}
		}
		dp[i] = tmp
	}
	//l为纵轴，子串长度
	//end为横轴，子串结束下标
	//
	for l := 1; l < len(s); l++ {
		for end := l; end < len(s); end++ {
			if s[end] != s[end-l] {
				dp[l][end] = false
			} else {
				if l < 2 {
					dp[l][end] = true
					y = l
					x = end
				} else if dp[l-2][end-1] {
					dp[l][end] = true
					y = l
					x = end
				}
			}
		}
	}
	res := s[x-y : x+1]
	return res
}