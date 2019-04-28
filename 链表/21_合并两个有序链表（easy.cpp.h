将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

//4-16
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
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        if(!l1&&l2)
            return l2;
        if(l1&&!l2)
            return l1;
        if(!l1&&!l2)
            return l1;
        	//所给链表存在头节点
		ListNode* res 	= l1;
        ListNode* head  = l1;
		
        if(l1->val>=l2->val){//设定新链表的头节点
            res = l2;
            l2 = l2->next;
        }
        else{
            res = l1;
            l1 = l1->next;
        }
       
        head = res;
       
		while(l1 && l2){
			if(l1->val>=l2->val){
				res->next = l2;
                l2=l2->next;
			}
		    else{
				res->next = l1;
                l1 = l1->next;
			}
            res = res->next;
		}
		
		if(l1)
			res->next = l1;
		else
			res->next = l2;
        
	
		return head;
    }
};




给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3



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
    ListNode* deleteDuplicates(ListNode* head) {
        
    }
};





