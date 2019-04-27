
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4

示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5

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
    ListNode* sortList(ListNode* head) {
		ListNode* in = head;
		ListNode* res = new ListNode(0);//新链表的头节点；
		ListNode* temp =new ListNode(0);
		temp = res;
		temp->next = NULL;
		ListNode* l = in;
		while(in){
            l = in->next;
			while(temp->next){
				if(temp->next->val > in->val){
					break;
				}
				temp = temp->next;
			}
			//新链表插入节点：
			in->next = temp->next;
			temp->next = in;
			temp = res;
			in = l;
		}
		return res->next;
    }
};


输入: 4->2->1->3
输出: 1->2->3->4