func repeatedNTimes(nums []int) int {
    m := make(map[int]int)
    n := len(nums) / 2
    for _, v := range nums {
        m[v] += 1
    }
    result := -1
    for k := range m {
        if m[k] == n {
            result = k
            break
        }
    }
    return result
}
