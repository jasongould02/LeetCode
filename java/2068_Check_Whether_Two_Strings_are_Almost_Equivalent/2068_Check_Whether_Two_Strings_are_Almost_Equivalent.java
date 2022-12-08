class Solution {
    public boolean checkAlmostEquivalent(String word1, String word2) {
        int[] letters = new int[26];
        for (char c : word1.toCharArray()) {
            letters[c-'a'] += 1;
        }
        for (char c : word2.toCharArray()) {
            letters[c-'a'] -= 1;
        }
        
        for (int i : letters) {
            if (Math.abs(i) > 3) {
                return false;
            }
        }
        return true;
    }
}
