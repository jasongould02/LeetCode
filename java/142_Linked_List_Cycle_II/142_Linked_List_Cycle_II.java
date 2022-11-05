/**
 * Definition for singly-linked list.
 * class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    public ListNode detectCycle(ListNode head) {
        HashSet<ListNode> set = new HashSet<ListNode>();
        ListNode cur = head;
        int pos = 0;
        while(cur != null) {
            if (!set.contains(cur.next)) {
                set.add(cur);
                cur = cur.next;
                pos++;
            } else {
                break;
            }
        }
        return cur != null ? cur.next : cur;
    }
}
