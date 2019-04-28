给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

示例 1:

输入: "Let's take LeetCode contest"
输出: "s'teL ekat edoCteeL tsetnoc" 
注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

func reverseWords(s string) string {
    b :=[] byte(s)
    
  //  k:=0    //从左往右一直遍历
    r:=0    //记录单词的右边界
    l:=0    //记录单词的左边界

    for k:=0;k<len(b);k++{
        if b[k]==32{
            r=k
            reverseWord(b[l:r])
           // j:=string(p)
          //  fmt.Println(j)
            l=k+1
        }   
    } 
    reverseWord(b[l:len(b)])//对最后一个单词进行反转
    s = string(b)
    return s
}

func reverseWord(s [] byte){
     r:=0
    //fmt.Printf("%d",l)
    for l:=len(s)-1;r<l ;l--{
            s[r],s[l]=s[l],s[r]
            r++
    }
}