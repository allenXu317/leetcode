
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]  
   1
    \
     2
    /
   3 

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    var stack []*TreeNode
    var res []int
    
    for root!=nil||len(stack)!=0{
        for root!=nil{
            stack = append(stack,root)  //父节点入栈;
            res = append(res,root.Val)  //前序，开始读父节点
            root = root.Left    //左节点入栈;
        }
        if len(stack)!=0{
            root = stack[len(stack)-1]
            root = root.Right
            stack = stack[0:len(stack)-1]
        }
    }
    return res
}