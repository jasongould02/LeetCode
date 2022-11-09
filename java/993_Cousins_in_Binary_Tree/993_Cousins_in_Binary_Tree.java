/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    TreeNode xParent = null;
    TreeNode yParent = null;
    int xDepth = -1;
    int yDepth = -1;
    public boolean isCousins(TreeNode root, int x, int y) {
        getDepth(root, x, y, 0);
        if (xParent != yParent && xDepth == yDepth) {
            return true;
        }
        return false;        
    }
    
    private void getDepth(TreeNode root, int x, int y, int depth) {
        if (root != null) {
            if (root.left != null) {
                if (root.left.val == x) {
                    xParent = root;
                    xDepth = depth;
                } else if (root.left.val == y) {
                    yParent = root;
                    yDepth = depth;
                }
            } 
            if (root.right != null) {
                if (root.right.val == x) {
                    xParent = root;
                    xDepth = depth;
                } else if (root.right.val == y) {
                    yParent = root;
                    yDepth = depth;
                }
            }
            
            getDepth(root.left, x, y, depth + 1);
            getDepth(root.right, x, y, depth +1);
            
        }
    }
}
