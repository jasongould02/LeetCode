/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    length := nodeCount(head)
    middle := length/2
    selectSecond := false
    
    if length & 1 == 1 {
        selectSecond = true
    }
    pos := 1
    cur := head
    for pos < middle {
        cur = cur.Next
        pos++
    }
    if selectSecond {
        return cur.Next
    } else {
        return cur
    }
}

func nodeCount(head *ListNode) int {
    cur := head
    total := 1
    if cur == nil {
        total = 0
    }
    for cur != nil {
        cur = cur.Next
        total++
    }
    return total
}
