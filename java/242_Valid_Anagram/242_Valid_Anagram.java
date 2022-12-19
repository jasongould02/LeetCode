class Solution {
    public boolean isAnagram(String s, String t) {
        int[] ss = new int[26];
        if (s.length() != t.length()) {
            return false;
        }
        for (int i = 0; i < s.length(); i++) {
            ss[s.charAt(i)-'a'] += 1;
            ss[t.charAt(i)-'a'] -= 1;
        }
        for (int i : ss) {
            if (i != 0) {
                return false;
            }
        }
        return true;
    }
}
