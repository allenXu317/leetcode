
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var stack []*TreeNode
    var res [] int
    
    for root!=nil||len(stack)!=0{//len(stack)如果不为0，则还有一个节点的值需要压栈
        for root!=nil{
            stack = append(stack,root)
           // res = append(res,root.Val)
            root = root.Left
        }
        if len(stack)!=0{
            root = stack[len(stack)-1]
            stack = stack[0:len(stack)-1]
            res = append(res,root.Val)
            root = root.Right
        }
    }
    return res
}