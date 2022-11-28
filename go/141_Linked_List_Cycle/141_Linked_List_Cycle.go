/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    m := make(map[*ListNode]bool)
    pos := head
    good := false
    for good == false {
        if pos == nil {
            break
        }
        good = pass(m, pos)
        pos = pos.Next
    }
    return good
}

func pass(m map[*ListNode]bool, pos *ListNode) bool {
    _, exists := m[pos]
    if exists {
        return true
    } else {
        m[pos] = true
        return false
    }
}
