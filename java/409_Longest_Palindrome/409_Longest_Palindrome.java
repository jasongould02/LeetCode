class Solution {
    public int longestPalindrome(String s) {
        int[] c = new int[128];
        
        for (char a : s.toCharArray()) {
            c[a] += 1;
        }

        int result = 0;
        for (int v : c) {
            result += v / 2 * 2;
            if ((result & 1) == 0 && (v & 1) == 1) {
                result += 1;
            }
        }
        return result;
    }
}
