class Solution {
    public int sumOfUnique(int[] nums) {
        HashMap<Integer, Integer> map = new HashMap<Integer, Integer>();
        for (int i : nums) {
            map.merge(i, 1, Integer::sum);
        }
        int total = 0;
        for (Integer key : map.keySet()) {
            if (map.get(key) == 1) {
                total += key;
            }
        }
        return total;
    }
}
