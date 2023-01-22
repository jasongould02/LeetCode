class Solution {
    public int minPartitions(String n) {
        char max = 0;
        for (char c : n.toCharArray()) {
            if (c == '9') {
                return 9;
            }
            if (c > max) {
                max = c;
            }
        }
        return (int) (max - '0');
    }
}
