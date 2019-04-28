你将得到一个字符串数组 A。

如果经过任意次数的移动，S == T，那么两个字符串 S 和 T 是特殊等价的。

 

一次移动包括选择两个索引 i 和 j，且 i ％ 2 == j ％ 2，并且交换 S[j] 和 S [i]。

现在规定，A 中的特殊等价字符串组是 A 的非空子集 S，这样不在 S 中的任何字符串与 S 中的任何字符串都不是特殊等价的。

 

返回 A 中特殊等价字符串组的数量。

 

示例 1：

输入：["a","b","c","a","c","c"]
输出：3
解释：3 组 ["a","a"]，["b"]，["c","c","c"]
示例 2：

输入：["aa","bb","ab","ba"]
输出：4
解释：4 组 ["aa"]，["bb"]，["ab"]，["ba"]
示例 3：

输入：["abc","acb","bac","bca","cab","cba"]
输出：3
解释：3 组 ["abc","cba"]，["acb","bca"]，["bac","cab"]
示例 4：

输入：["abcd","cdab","adcb","cbad"]
输出：1
解释：1 组 ["abcd","cdab","adcb","cbad"]
 

提示：

1 <= A.length <= 1000
1 <= A[i].length <= 20
所有 A[i] 都具有相同的长度。
所有 A[i] 都只由小写字母组成。



//判断奇数位或者偶数位上的字符是否相等：

/*
题目太晦涩。 根据移动规则来说，就是如果两个字符串，奇数位的字母一样，偶数位的字母一样，这样的字符串相互等价。 根据示例，能看出来，任何字符串和自身也算是等价的。

所以就是最后看有几个归一化的字符串，这个数就是结果

长度小于等于2的字符串，就只能是自身与自身等价。这样的测试用例，直接set一下，返回长度即可。
每一个字符串取偶数位排序，奇数位排序，合并成一个字符串（怎么样合并无所谓，交叉，前后颠倒都可以），放入新数组
返回新数组set后的长度。
*/ 

//借鉴了讨论中一位同学的意见



func numSpecialEquivGroups(A []string) int {
 
    
    //找是否存在不移动的元素
    
    num := len(A)
    length := len(A[0])
    
    //直接两两比较
    //分别取出奇数位和偶数位的字符：
  //  var evenSet  map[int] string
 //   var oddSet   map[int] string
    
    var ResStrings []string
    
    for j:=0;j<num;j++{
        var evenString []byte
        var oddString []byte
        p := []byte(A[j])
        for i:=0;i<length;i++{
            if i%2==0{//奇数位：
                oddString = append(oddString,p[i])
            }else{//偶数位
                evenString = append(evenString,p[i])
            } 
         }
         evenstring := string(evenString[:])//偶数位字符串
         oddstring := string(oddString[:])//奇数位字符串
         evenstring = sortString(evenstring)
         oddstring = sortString(oddstring)
         resString := oddstring+evenstring
        // fmt.Printf("偶数位是：%s\n",evenstring)
        // fmt.Printf("奇数位是：%s\n",oddstring)
       //  fmt.Printf("最后的组合是：%s\n",resString)
         ResStrings = append( ResStrings,resString)//放入最后的字符串组合数组；
    }
    
    //计算最后的组合数：
    for i:=0;i<num;i++{
        for j:=i+1;j<len(ResStrings);{
            if ResStrings[i]==ResStrings[j] {
               // fmt.Printf("删除了:%s\n",ResStrings[j]);
                ResStrings = append(ResStrings[:j],ResStrings[j+1:]...)//将这个相同的从组合数组中删除
                continue
            }else{
                j++ //删除时对于j的控制
            }
        }
    }
    count := 0
    //还剩最后一个在数组中，排查：
    for j:=0;j<len(ResStrings);j++ {
        if (ResStrings[j] == ResStrings[len(ResStrings)-1] && j!=(len(ResStrings)-1)){
           count++
        }
    }
    return len(ResStrings)-count
}

//给字符串排序的方法：
func sortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}



