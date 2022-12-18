class Solution {
    public int uniqueMorseRepresentations(String[] words) {
        String[] set = new String[] {".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."};
        
        HashSet<String> results = new HashSet<String>();
        int count = 0;
        for (String w : words) {
            String cur = "";
            for (char c : w.toCharArray()) {
                cur += set[c - 'a'];
            }
            if (results.add(cur)) {
                count += 1;
            }
        }
        return count;
    }
}
