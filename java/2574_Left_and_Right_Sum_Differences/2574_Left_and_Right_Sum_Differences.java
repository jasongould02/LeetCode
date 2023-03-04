class Solution {
    public int[] leftRigthDifference(int[] nums) {
        int[] res = new int[nums.length];
        int total = 0;
        for (int i : nums) {
            total += i;
        }
        int lsum = 0;
        for (int i = 0; i < nums.length; i++) {
            total -= nums[i];
            res[i] = Math.abs(lsum - total);
            lsum += nums[i];
        }
        return res;
    }
}
