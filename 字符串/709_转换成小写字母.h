709. 转换成小写字母

实现函数 ToLowerCase()，该函数接收一个字符串参数 str，并将该字符串中的大写字母转换成小写字母，之后返回新的字符串。

 

示例 1：

输入: "Hello"
输出: "hello"
示例 2：

输入: "here"
输出: "here"
示例 3：

输入: "LOVELY"
输出: "lovely"

//2019.4.22

//ASCII码：
//大写A：65     Z：90
//小写a：97		z：122


class Solution {
public:
    string toLowerCase(string str) {
        string res;
		
		for(auto &i : str){
			if(i>='A'&&i<='Z'){
				i = i+32;
			}
		}
		return str;
    }
};