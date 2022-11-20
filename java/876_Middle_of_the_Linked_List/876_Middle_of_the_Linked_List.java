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
    public ListNode middleNode(ListNode head) {
        int length = nodeCount(head);
        int middle = length/2;
        boolean selectSecond = length % 2 == 0 ? false : true;
        
        int pos = 1;
        ListNode curr = head;
        while (pos < middle) {
            curr = curr.next;
            pos++;
        }
        return selectSecond ? curr.next : curr;
    }
    
    private int nodeCount(ListNode head) {
        ListNode curr = head;
        int total = curr == null ? 0 : 1;
        while(curr != null) {
            curr = curr.next;
            total++;
        }
        return total;
    }
}
