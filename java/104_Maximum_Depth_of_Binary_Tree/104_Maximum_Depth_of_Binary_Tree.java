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
    public int maxDepth(TreeNode root) {
        return checkBranch(root);
    }
    
    public int max(int a, int b) {
        return a > b ? a : b;
    }
    
    private int checkBranch(TreeNode root) {
        if (root == null) {return 0;}
        return max(checkBranch(root.left), checkBranch(root.right)) + 1;
        
    }
}