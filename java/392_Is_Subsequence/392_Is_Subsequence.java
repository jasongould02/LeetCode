class Solution {
    public boolean isSubsequence(String s, String t) {
        char[] seq = s.toCharArray();
        int spos = 0;

        if ((t.length() == 0 && s.length() == 0) || s.length() == 0) {
            return true;
        } else if (t.length() == 0 || s.length() > t.length()) {
            return false;
        }
        
        for (char c : t.toCharArray()) {
            if (c == seq[spos]) {
                spos++;
                if (spos >= seq.length) {
                    return true;
                }
            }
        }
        return false;
    }
}
