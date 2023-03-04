class Solution {
    public int mostWordsFound(String[] sentences) {
        int max = 0;
        for (String line : sentences) {
            int count = line.length() - line.replace(" ", "").  length();
            if (count > max) {
                max = count;
            }
        }
        return max+1;
    }

