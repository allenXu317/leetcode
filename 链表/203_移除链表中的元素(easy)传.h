删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* removeElements(ListNode* head, int val) {
        if(!head)
            return head;
		if(!head->next && (head->val==val))//只有一个节点且需要删除时
			return NULL;
		else if(!head->next && (head->val!=val))
			return head;//不需要删除时；
			
		
        ListNode* res = head;
   
		while(res){//节点数大于1个时；
			if((res->val==val) && (res->val == head->val)){//需要删除的节点为第一个节点时：
				head = head->next;
				res = head;
				continue;//新链表开始下一次寻找；
			}
			if(res->next && res->next->val == val){
				res->next = res->next->next;
			}
			if(res->next && res->next->val==val)//不需要连续删除
				continue;
			else
				res = res->next;
		}
		
		return head;
    }
};
[1,2,2,1]
2

[1,2,6,3,4,5,6]
6