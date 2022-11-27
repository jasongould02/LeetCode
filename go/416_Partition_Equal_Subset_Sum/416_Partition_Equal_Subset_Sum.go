func canPartition(nums []int) bool {
    totalSum := 0
    for _, n := range nums {
        totalSum += n
    }

    if totalSum & 1 == 1 {
        return false
    }

    totalSum /= 2
    dp := make([]bool, totalSum + 1)
    dp[0] = true
    for i := 1; i < len(nums); i++ {
        for j := totalSum; j >= nums[i-1]; j-- {
            dp[j] = dp[j] || dp[j - nums[i-1]]
        }
    }
    return dp[totalSum]
}
