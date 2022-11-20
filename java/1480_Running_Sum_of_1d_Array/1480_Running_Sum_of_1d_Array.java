class Solution {
    public int[] runningSum(int[] nums) {
        int total = nums[0];
        int temp = 0;
        for (int i = 1; i < nums.length; i++) {
            temp = nums[i];
            nums[i] = total + nums[i];
            total += temp;
        }
        return nums;
    }
}
