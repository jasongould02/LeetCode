class Solution {
    public int maximumDifference(int[] nums) {
        int res = -1;
        for (int i = 0; i < nums.length-1; i++) {
            for (int j = i+1; j < nums.length; j++) {
                if (nums[j] - nums[i] > res && nums[j] - nums[i]  != 0) {
                    res = nums[j] - nums[i];
                }
            }
        }
        return res;
    }
}
