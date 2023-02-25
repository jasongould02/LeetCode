class Solution {
    public int minEatingSpeed(int[] piles, int h) {
        int left = 1;
        int right = 0;

        for (int i : piles) {
            if (i > right) right = i;
        }
        
        while (left < right) {
            int mid = left + (right - left) / 2;
            int sum = 0;

            for (int i : piles) {
                sum += i / mid;
                if (i % mid != 0) {
                    sum++;
                }
            }

            if (sum <= h) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }
        
        return left;
    }
}
