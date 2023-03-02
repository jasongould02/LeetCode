func distributeCandies(candyType []int) int {
    m := make(map[int]int)
    n := len(candyType)
    unique := 0
    for _, v := range candyType {
        _, exists := m[v]
        if !exists {
            unique++
        }
        m[v] += 1
    }
    if unique < n/2 {
        return unique
    }
    return n / 2
}
