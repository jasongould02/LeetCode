class Solution {
    public int singleNumber(int[] nums) {
        HashMap<Integer, Integer> map = new HashMap<Integer, Integer>();
        
        for (int n : nums) {
            if (map.containsKey(n)) {
                map.remove(n);
            } else {
                map.merge(n, 1, Integer::sum);
            }
        }
        for (int key : map.keySet()) {
            return key;
        }
        return 0;
    }
}
