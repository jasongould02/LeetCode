/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class P160 {
    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        int alen = length(headA);
        int blen = length(headB);
        
        int min, diff;
        boolean ashorter;
        
        if (alen > blen) {
            diff = alen - blen;
            ashorter = false;
        } else {
            diff = blen - alen;
            ashorter = true;
        }
        ListNode a = headA;
        ListNode b = headB;
        
        while (diff > 0) {
            if (ashorter) {
                b = b.next;
            } else {
                a = a.next;
            }
            diff--;
        }
        
        while (a != b) {
            a = a.next;
            b = b.next;
        }
        
        return a;
    }
    
    private int length(ListNode head) {
        int length = 0;
        while(head != null) {
            head = head.next;
            length++;
        }
        return length;
    }
}