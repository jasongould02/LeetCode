func singleNumber(nums []int) int {
    m := make(map[int]int)

    for _, n := range nums {
        _, exists := m[n]
        if exists {
            delete(m, n)
        } else {
            m[n] += 1
        }
    }
    for k := range m {
        return k
    }
    return 0
}
