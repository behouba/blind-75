/*
FROM: https://leetcode.com/problems/add-two-numbers/

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

*/

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
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *n1 = l1;
        ListNode *n2 = l2;
        
        int rest = 0;
        ListNode *sum = NULL;
        ListNode *nSum;
        while(n1 != NULL || n2 != NULL) {
            int val1 = n1 != NULL ? n1->val : 0;
            int val2 = n2 != NULL ? n2->val : 0;
            int s = val1 + val2 + rest;
            ListNode *node;
            if(s > 9) {
                rest = 1; 
                node = new ListNode(s-10);
            } else {
                rest = 0;
                node = new ListNode(s);
            }
            if (sum == NULL) {
                sum = node;
                nSum = sum;
            } else {
                nSum->next = node;
                nSum = nSum->next;
            }
            
            if(n1 != NULL) {
                n1 = n1->next;
            }
            if(n2 != NULL) {
                n2 = n2->next;
            }
            
        }
        if(rest) {
            nSum->next = new ListNode(rest);
        }
        return sum;
    }
};
