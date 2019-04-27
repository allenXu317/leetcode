给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3

2019/4/17
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
       if(!head)
            return head;
       int temp = head->val;
		ListNode* ergo = head;//ergodic
		ListNode* left = head;
		ListNode* right = left;
		ergo = ergo->next;//ergo往后移一位；
		
		while(ergo){
			if(temp!=ergo->val){
				left = ergo;
				temp = ergo->val;
				ergo = ergo->next;
			}
			else{
				left->next = ergo->next;
				ergo = ergo->next;
			}
		}
		return head;
    }
};



