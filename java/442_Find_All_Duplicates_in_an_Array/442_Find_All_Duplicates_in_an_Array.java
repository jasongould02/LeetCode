class Solution {
    public List<Integer> findDuplicates(int[] nums) {
        if (nums.length == 1) {
            return new ArrayList<Integer>();
        }
        Arrays.sort(nums);
        ArrayList<Integer> result = new ArrayList<Integer>();
        for (int i = 0; i < nums.length - 1; i++) {
            if (nums[i] == nums[i+1]) {
                result.add(nums[i]);
            }
        }
        return result;
    }
}
