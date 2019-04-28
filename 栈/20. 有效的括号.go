给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true



func isValid(s string) bool {
    //用slice实现一个栈
    
    k := strings.Split(s,"")
    if len(k)%2!=0{
        return false
    }
    
    var stack []string
    
    
    for _, p := range k{
        //右括号入栈；
        if (p=="("||p=="{"||p=="["){
            //入栈：
            stack = append(stack,p)
        }else{//左括号，栈顶元素出栈，与左括号比较：
            if len(stack)==0{
                return false;
            }
            t := stack[len(stack)-1]//栈顶元素；
            if (t!="("&&p==")"){
                return false
            }else if (t!="["&&p=="]"){
                return false
            }else if (t!="{"&&p=="}"){
                return false
            }
            stack = stack[:len(stack)-1] //栈顶弹出，压缩栈
        }
    }
    if len(stack)==0{
        return true
    }
    return false
}