给定一个由空格分割单词的句子 S。每个单词只包含大写或小写字母。

我们要将句子转换为 “Goat Latin”（一种类似于 猪拉丁文 - Pig Latin 的虚构语言）。

山羊拉丁文的规则如下：

如果单词以元音开头（a, e, i, o, u），在单词后添加"ma"。
例如，单词"apple"变为"applema"。

如果单词以辅音字母开头（即非元音字母），移除第一个字符并将它放到末尾，之后再添加"ma"。
例如，单词"goat"变为"oatgma"。

根据单词在句子中的索引，在单词最后添加与索引相同数量的字母'a'，索引从1开始。
例如，在第一个单词后添加"a"，在第二个单词后添加"aa"，以此类推。
返回将 S 转换为山羊拉丁文后的句子。

示例 1:

输入: "I speak Goat Latin"
输出: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
示例 2:

输入: "The quick brown fox jumped over the lazy dog"
输出: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
说明:

S 中仅包含大小写字母和空格。单词间有且仅有一个空格。
1 <= S.length <= 150。









func toGoatLatin(S string) string {
    
    s := strings.Split(S," ") //把每个单词分割出来；
    
    vowel := [...]string{"a", "e", "i", "o", "u","A","E","I","O","U"}   //元音字母查询
    
    var c [] []string
    
    for _,p := range s{
        t := strings.Split(p,"")    //取出每个字母
        c = append(c,t) //加入到c中
    }
    
    //测试：---通过
    /*
    for _,p := range c{
        fmt.Printf("字母数组：%s\n",p)
        for _,o := range p{
            fmt.Printf("字母：%s\n",o)
        }
    }
    */
    
   // var flag bool
    i := 0
    
    var res []string
    
    
    for _,p := range c{
        i++
        k := goatLatin(vowel[0:],p[0:],i)
        temp := strings.Join(k,"")
        res = append(res,temp)
       // fmt.Printf("转换后：%s\n",k)
       // fmt.Printf("转换后：%s\n",res)
    }
    return strings.Join(res," ")
}

//判断元素是否存在于数组中；
func inArray(s []string, i string) bool{
    
    for _,p :=range s{
        if i == p{
            return true
        }
    }
    return false
}

//将单词改为GL；
//p传入为每个单词；
//应当为c
func goatLatin(a []string, p []string, i int) []string{
    var gl [] string
    var flag bool
        
        flag = inArray(a[0:],p[0])//判断首字母；---a为元音字母数组；
        
        if flag{ //此时为元音;在后面直接加ma；
            l := strings.Join(p[0:],"")
            gl = append(gl,l)
        }else{
            //首字母放到最后;
            l := strings.Join(p[1:],"")
            gl = append(gl,l)
            gl = append(gl,p[0])
        }
    
        gl = append(gl,"ma")    //末尾加上ma
    
    //往末尾添加a：
    for j:=0;j<i;j++{
        gl = append(gl,"a")
    }
    
    return gl
}

