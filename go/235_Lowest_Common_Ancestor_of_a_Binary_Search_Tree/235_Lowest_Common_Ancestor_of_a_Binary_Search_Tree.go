/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    for (root.Val - p.Val) * (root.Val - q.Val) > 0 {
        if p.Val < root.Val {
            root = root.Left
        } else if p.Val > root.Val {
            root = root.Right
        } else {
            break;
        }
    }
    return root
}
