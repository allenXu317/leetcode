//创建一个队列
type Quee struct {
    count int
    head int
    tail int
    list []int
}

func NewQuee() *Quee {
    q := new(Quee)
    q.count = 0
    q.head = 0
    q.tail = 0
    q.list = make([]int,0)
    return q
}
//入列
func (q *Quee) Push(v int) {
    q.count++
    q.tail++
    q.list = append(q.list,v)
}
//出列
func (q *Quee) Pull() {
    if q.head >= q.tail{
        fmt.Println("队列已空")
    }
    q.head++
    q.count--
    q.list = q.list[q.head:]
    q.head = 0
}
//获得队头元素
func (q *Quee) GetHead () int {
    res := q.list[0]
    return res
}

func findContinuousSequence(target int) [][]int {
    //缩小查找区域
    res := make([][]int,0)
    mid := (target+1)/2 
    q := NewQuee()
    sum := 0
    for i:=1;i<=mid;i++ {
        q.Push(i)
        sum += i
        for sum > target {
            //队头出列
            sum -= q.GetHead()
            q.Pull()
        }
        if sum == target {
            //将队列中的元素存入数组中
            //队头出列
            //sum减去队头
            tmp := make([]int,0)
            for _,v := range q.list {
                tmp = append(tmp,v)
            }
            res = append(res,tmp)
            sum -= q.GetHead()
            q.Pull()
        } 
    }
    return res
}