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
    public ListNode sortList(ListNode head) {
        if (head == null) {return null;}
        ArrayList<Integer> list = new ArrayList<Integer>();
        while (head != null) {
            list.add(head.val);
            head = head.next;
        }
        Collections.sort(list, Collections.reverseOrder());
        ListNode cur = new ListNode();
        for (Integer g : list) {
            ListNode temp = new ListNode(g, cur.next);
            cur.next = temp;
        }
        return cur.next;
    }
}
