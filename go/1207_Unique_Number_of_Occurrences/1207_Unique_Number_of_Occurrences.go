func uniqueOccurrences(arr []int) bool {
    m := make(map[int]int)
    counts := make(map[int]bool)
    for _, v := range arr {
        m[v] += 1
    }
    for _, v := range m {
        _, exists := counts[v]
        if exists {
            return false
        } else {
            counts[v] = true
        }
    }
    return true
}
