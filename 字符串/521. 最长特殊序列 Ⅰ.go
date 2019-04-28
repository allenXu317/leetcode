给定两个字符串，你需要从这两个字符串中找出最长的特殊序列。最长特殊序列定义如下：该序列为某字符串独有的最长子序列（即不能是其他字符串的子序列）。

子序列可以通过删去字符串中的某些字符实现，但不能改变剩余字符的相对顺序。空序列为所有字符串的子序列，任何字符串为其自身的子序列。

输入为两个字符串，输出最长特殊序列的长度。如果不存在，则返回 -1。

示例 :

输入: "aba", "cdc"
输出: 3
解析: 最长特殊序列可为 "aba" (或 "cdc")

func findLUSlength(a string, b string) int {
 //可操作内存空间;
    l := []byte(a)
    m := []byte(b)
    step := 0
    if len(a)==0&&len(b)==0{
        return -1
    }
    if len(a)>len(b){
        step = len(a)
    }else{
        step = len(b)
    }
    if len(a)==len(b) && l[0]==m[0]{
        step = -1
    }
    return step;
}