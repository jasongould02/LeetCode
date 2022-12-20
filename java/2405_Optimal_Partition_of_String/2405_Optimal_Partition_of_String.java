class Solution {
    public int partitionString(String s) {
        int count = 1;
        int[] set = new int[26];
        for (char c : s.toCharArray()) {
            if (set[c-'a'] == 1) {
                count += 1;
                set = new int[26];
            }
            set[c-'a'] += 1;
        }
        return count;
    }
}
