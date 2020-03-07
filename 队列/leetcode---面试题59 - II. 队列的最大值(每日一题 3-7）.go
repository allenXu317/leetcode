请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：

输入: 
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]
示例 2：

输入: 
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]
 

限制：

1 <= push_back,pop_front,max_value的总操作数 <= 10000
1 <= value <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


我的博客：https://blog.csdn.net/XUCHEN1230/article/details/104712501

type MaxQueue struct {
    head int
    tail int
    count int
    max int
    list []int
}
 
 
func Constructor() MaxQueue {
    q := MaxQueue{
        head : 0,
        tail : 0,
        count :0,
        max : 0,
        list : make([]int ,0),
    }
    return q
}
 
 
func (this *MaxQueue) Max_value() int {
    if this.count == 0 {
        return -1
    }
    // fmt.Println("调用Max_Value:",this.max)
    return this.max
}
 
 
func (this *MaxQueue) Push_back(value int)  {
    //更新max
    if this.count == 0 {
        this.max = value
    } else if  value > this.max {
            this.max = value
    } 
    this.list = append(this.list,value)
    this.tail++
    this.count++
    // fmt.Println(this.max,this.list)
    // fmt.Println("调用Push_Back"," ","value: ",value," ",this.list," ",this.max)
}
 
 
func (this *MaxQueue) Pop_front() int {
    if this.count == 0{
        return -1
    }
    pop := this.list[0]
    this.count--
    this.head++
    this.list = this.list[this.head:]
    this.head = 0
    //需要更新max
    if pop == this.max && this.count > 0{
        tmp := this.list[len(this.list)-1]
        for _,v := range this.list {
            if v > tmp {
                tmp = v
            }
        }
        this.max = tmp
    } else if this.count == 0{
        this.max = 0
    }
    // fmt.Println("调用pop"," ",this.list," ",this.max," ",pop)
    return pop
}
 
 
 
/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */