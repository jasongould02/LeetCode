func numIdenticalPairs(nums []int) int {
    n := make([]int, 101)

    for i := 0; i < len(nums); i++ {
        n[nums[i]] += 1
    }
    total := 0
    for i := 0; i < len(n); i++ {
        total += (n[i] * (n[i]-1))/2
    }
    return total
}
