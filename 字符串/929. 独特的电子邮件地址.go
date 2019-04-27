
//go踩坑：
/*
go的import机制：
1.首先设置环境变量：
对$GOPATH进行设置，这个路径下的src就是导入包时寻找的路径；

2.在import包时，是寻找的文件夹路径：
例如：
import "fmt"
go自动在$GOROOT的src中寻找 名为 fmt 的文件夹，所以在导入自定义包时，应该在$GOPATH的src下新建一个同名文件夹才能导入；

*/

/*
go强制类型转换：
int(变量名)
*/
func numUniqueEmails(emails []string) int {
  //  i :=0
    var res []string
    var flag bool = false
    var m int = 0
    for _,s:=range emails{
        k:=emailStr(s)
      //  fmt.Println(k)
        
        for _,t:=range res{
           // fmt.Println(t)
            if k==t{
                flag = true
                break;
            }
        }
        if !flag{
            res = append(res,k)
            m++
        }
        flag = false;
    }
    return m
}

func emailStr (email string) string {
    //对email进行处理：
    var flag bool = false;
    var flag2 bool = false;
    var res []byte;
    
    s :=[]byte(email)   //分配一个新的字节数组保存email的内容；
    
    for _,c:=range s{
        //fmt.Printf("%c",c);//测试通过
        
        if c=='+'{
            flag = true;
        }
        if c=='@'{//进入到域名中：
            flag2 = true;
        }
        
        if flag2{
            res = append(res,c)
            continue
        }
        
        if !flag{//在本地名称中进行处理;
            if c!='.'{//直接插入到res中
                res = append(res,c)
            }else if c=='.'{
                continue
            }
        }
    }
    
    out := string(res)
    
    return out
}