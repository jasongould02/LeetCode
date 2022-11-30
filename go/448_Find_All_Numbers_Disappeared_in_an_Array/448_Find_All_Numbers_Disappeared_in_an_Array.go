func findDisappearedNumbers(nums []int) []int {
    result := make([]int, len(nums))
    for _, n := range nums {
        result[n-1] += 1
    }
    out := make([]int, 0)
    for i, n := range result {
        if n == 0 {
            out = append(out, i+1)
        }
    }
    return out
}
