class Solution {
    public int[] arrayChange(int[] nums, int[][] operations) {
        HashMap<Integer, Integer> map = new HashMap<Integer, Integer>();
        for (int i = 0; i < nums.length; i++) {
            map.put(nums[i], i);
        }
        for (int i = 0; i < operations.length; i++) {
            int origpos = map.get(operations[i][0]);
            map.remove(operations[i][0]);
            nums[origpos] = operations[i][1];
            map.put(operations[i][1], origpos);
        }
        return nums;
    }
}
