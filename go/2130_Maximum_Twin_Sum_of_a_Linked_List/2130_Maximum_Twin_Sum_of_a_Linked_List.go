/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    cur := head
    n := 0
    list := make([]int, 0, 50)
    for cur != nil {
        n++
        list = append(list, cur.Val)
        cur = cur.Next
    }
    max := 0
    for i, v := range list {
        twinIndex := n-1-i
        if i <= (n/2)-1 {
            if list[twinIndex] + v > max {
                max = list[twinIndex] + v
            }
        }
    }
    return max
}
