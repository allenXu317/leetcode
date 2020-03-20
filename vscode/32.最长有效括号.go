/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {
	type Quee struct {
		head int
		tail int
		list []rune
	}

	// func NewQuee() *Quee {
	// 	q := new(Quee)
	// 	q.head = 0
	// 	q.tail = 0
	// 	q.list = make([]rune,0)
	// 	return q
	// }

	func (q *Quee) Push(i rune) {
		q.list = append(q.list,i)
		tail++
	}

	func (q *Quee) Pop() rune{
		if len(q.list) == 0{
			return 'x'
		}
		r := q.list[0]
		q.list = q.list[1:]
		tail--
		return r
	}
	
	res := Run(s)

	return res
}
func Run (s string) run {
	dp := make([]int, len(s))

	q := new(Quee)
	q.head = 0
	q.tail = 0
	q.list = make([]rune,0)

	fmt.Println(q)

	dp[0] = 0
	count := 0
	res := 0

	for i := 1; i < len(s); i++ {
		if s[i-1] == '(' && s[i] == ')' {
			count += 2
			dp[i] = count + dp[i-1]
		} else {
			dp[i] = 0
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

// @lc code=end

