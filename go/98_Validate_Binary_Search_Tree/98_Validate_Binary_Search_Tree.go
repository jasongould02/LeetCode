/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "math"
func isValidBST(root *TreeNode) bool {
    return check(root, math.MinInt64, math.MaxInt64)
}

func check(root *TreeNode, lower, upper int64) bool { 
    if root == nil {
        return true
    }
    t := int64(root.Val)
    if (t <= lower) || (t >= upper) {
        return false
    }
    if !check(root.Left, lower, t) || !check(root.Right, t, upper) {
        return false
    }
    return true
}