class Solution {
    public int maxSum(int[][] grid) {
        int max = 0;
        for (int i = 0; i < grid.length-2; i++) {
            for (int j = 0; j < grid[0].length-2; j++) {
                int r1 = grid[i][j] + grid[i][j+1] + grid[i][j+2];
                int r2 = grid[i+1][j+1];
                int r3 = grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2];
                max = Math.max(r1 + r2 + r3, max);
            }
        }
        return max;
    }
}
