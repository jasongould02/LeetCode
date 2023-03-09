func finalPrices(prices []int) []int {
    ans := make([]int, len(prices))
    copy(ans, prices)
    for i := 0; i < len(prices)-1; i++ {
        for j := i+1; j < len(prices); j++ {
            if prices[j] <= prices[i] {
                ans[i] = prices[i] - prices[j]
                break
            }
        }
    }
    return ans
}
