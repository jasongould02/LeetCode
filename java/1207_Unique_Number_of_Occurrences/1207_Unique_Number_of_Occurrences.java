class Solution {
    public boolean uniqueOccurrences(int[] arr) {
        HashMap<Integer, Integer> map = new HashMap<Integer, Integer>();
        for (int v : arr) {
            map.merge(v, 1, Integer::sum);
        }
        HashSet<Integer> counts = new HashSet<Integer>();
        for (int key : map.values()) {
            if (counts.contains(key)) {
                return false;
            } else {
                counts.add(key);
            }
        }
        return true;
    }
}
