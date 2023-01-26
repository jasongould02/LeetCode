/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode mergeNodes(ListNode head) {
        ListNode newList = new ListNode();
        ListNode current = newList;
        int sum = 0;
        while(head != null) {
            if (head.val == 0 && sum != 0) {
                current.next = new ListNode();
                current = current.next;
                current.val = sum;
                sum = 0;
            }
            sum += head.val;
            head = head.next;
        }
        return newList.next;
    }
}
