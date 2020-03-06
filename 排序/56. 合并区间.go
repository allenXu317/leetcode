package leetcode

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	res := make([][]int, 0)
	//对 intervals 进行排序，按每个子数组第一个元素排序
	SortQuick(intervals)
	//对 intervals 进行二次排序，按每个子数组第二个元素排序
	res = SortMerge(intervals, len(intervals), 0, res)
	return res
}
func Exc(i int, j int, num [][]int) {
	num[i], num[j] = num[j], num[i]
}

//插入排序--二维数组
func SortInsert(num [][]int) {
	for i := 0; i < len(num)-1; i++ {
		t := i + 1
		if num[t][1] > num[i][1] {
			for l := 0; l <= i; l++ {
				if num[l][1] < num[t][1] {
					Exc(l, t, num)
				}
			}
		}
	}
}

//快速排序
func SortQuick(num [][]int) {
	//fmt.Println(num)
	length := len(num) - 1
	if length+1 <= 11 {
		SortInsert(num)
		return
	}
	var flag int
	if length+1 == 0 {
		return
	}
	a := num[0][1]
	x := (length + 1) / 2
	b := num[x][1]
	c := num[length][1]
	flag, mid := FindMid(a, b, c, x, length)
	num[flag], num[length] = num[length], num[flag]
	flag = length
	// fmt.Println(num)
	i := 0
	for j := 0; j <= length; j++ {
		//排序规则，从大到小排序，且i与j需要不相等，避免与下面得 num[i] 和 num[j]交换形成死循环
		if num[j][1] > mid {
			//遇到 i == flag 得情况说明分区点提前交换，此时更新key值
			num[i], num[j] = num[j], num[i]
			i++
			// fmt.Println(i,j,num)
		}
	}
	//防止不必要得交换

	num[i], num[flag] = num[flag], num[i]

	//递归调用
	// fmt.Println(mid,num)
	SortQuick(num[:i])

	SortQuick(num[i+1:])

}

//三点取中法选取分区点，此方法比固定取某个key点更加优化一些

//传出两个值，分别为key 和 对应得 value

//c语言可用数组传递

func FindMid(a, b, c, x, l int) (int, int) {
	// fmt.Println(a,b,c)
	max := a
	min := b
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	if a < min {
		min = a
	}
	if c < min {
		min = c
	}
	// fmt.Println(max,min)
	if a < max && a > min {
		return 0, a
	}
	if b < max && b > min {
		return x, b
	}
	if c < max && c > min {
		return l, c
	}
	return 0, a
}

// 得到重叠区间
func SortMerge(num [][]int, l int, i int, res [][]int) [][]int {
	if len(num) > 1 && len(res) == 0 {
		if num[0][0] >= num[1][0] && num[0][0] <= num[1][1] { //此时区间重叠
			num[1][1] = num[0][1]
			num[0][1] = 'a'
		} else if num[0][0] <= num[1][0] && num[0][1] >= num[1][1] {
			num[1][1] = num[0][1]
			num[1][0] = num[0][0]
			num[0][1] = 'a'
		}
		res = SortMerge(num[1:], l, i+1, res)
	}
	if num[0][1] != 'a' {
		res = append(res, num[0])
	}
	return res
}
