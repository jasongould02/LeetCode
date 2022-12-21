class Solution {
    public int addDigits(int num) {
        while (num > 9) {
            num = sumDigits(num);
        }
        return num;
    }
    
    public int sumDigits(int n) {
        int total = 0;
        while (n > 0) {
            total += n % 10;
            int digit = n % 10;
            n = n / 10;
        }
        return total;
    }
}
