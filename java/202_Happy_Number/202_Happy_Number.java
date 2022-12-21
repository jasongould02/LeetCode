class Solution {
    public boolean isHappy(int n) {
        HashMap<Integer, Boolean> map = new HashMap<Integer, Boolean>();
        while (n > 1) {
            if (map.containsKey(n)) {
                return false;
            }
            map.put(n, true);
            n = sumNum(n);
        }
        return true;
    }
    
    private int sumNum(int n) {
        int total = 0;
        while (n > 0) {
            int digit = n % 10;
            n = n / 10;
            total += digit * digit;
        }
        return total;
    }
}
