class Solution {
public:
    int countDistinctIntegers(vector<int>& nums) {
        set<int> numbers;
        int size = nums.size();
        int result = 0;
        for (int i = 0; i < size; i++) {
            int reverse = reverseNumber(nums[i]);
            if (numbers.count(nums[i]) < 1) {
                numbers.insert(nums[i]);
                result += 1;
            }
            if (numbers.count(reverse) < 1) {
                numbers.insert(reverse);
                result += 1;
            }
        }
        
        return result;
    }
    
    int reverseNumber(int number) {
        int reversed = 0;
        while(number > 0) {
            reversed = reversed*10 + number%10;
            number /= 10;
        }
        return reversed;
    }
};
