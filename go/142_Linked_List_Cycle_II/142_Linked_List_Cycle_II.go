/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    m := make(map[*ListNode]int)
    cur := head
    pos := 0
    for cur != nil {
        _, exists := m[cur]
        if !exists {
            m[cur] = cur.Val
            cur = cur.Next
            pos++
        } else {
            break
        }
    }
    return cur
}

