func findLonely(nums []int) []int {
    m := make(map[int]int)
    for _, v := range nums {
        m[v] += 1
    }
    result := make([]int, 0, 10)
    for k := range m {
        if m[k] == 1 {
            _, left := m[k-1]
            _, right := m[k+1]
            if (!left && !right) {
                result = append(result, k)
            }
        }
    }
    return result
}
