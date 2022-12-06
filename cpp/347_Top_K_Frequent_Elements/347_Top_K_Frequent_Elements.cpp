class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        map<int, int> m;
        for (int n : nums) {
            m[n] += 1;
        }
        
        vector<pair<int, int>> sortedNumbers;
        for (auto key : m) {
            sortedNumbers.push_back({key.second, key.first});
        }
        sort(sortedNumbers.rbegin(), sortedNumbers.rend());
        
        vector<int> result;
        for (int i = 0; i < k; i++) {
            result.push_back(sortedNumbers[i].second);
        }
        
        return result;
    }
};
