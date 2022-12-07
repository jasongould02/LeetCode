class Solution {
    public int repeatedNTimes(int[] nums) {
        int n = nums.length / 2;
        HashMap<Integer, Integer> map = new HashMap<Integer, Integer>();
        for (int i : nums) {
            map.merge(i, 1, Integer::sum);
        }

        int result = -1;
        for (int key : map.keySet()) {
            if (map.get(key) == n) {
                result = key;
                break;
            }
        }
        return result;
    }
}
