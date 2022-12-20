class Solution {
    public int sumOfUnique(int[] nums) {
        int[] values = new int[101];
        for (int i : nums) {
            values[i] += 1;
        }
        for (int i = 1; i < values.length; i++) {
            if (values[i] == 1) {
                values[0] += i;
            }
        }
        return values[0];
    }
}
