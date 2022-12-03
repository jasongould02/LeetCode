func findWinners(matches [][]int) [][]int {
    m := make(map[int]int)
    for _, n := range matches {
        m[n[1]] += 1
    }
    p1m := make(map[int]int)

    for _, n := range matches {
        _, exists := m[n[0]]
        if !exists {
            p1m[n[0]] = 1
        }
    }
    p1 := make([]int, 0, len(p1m))
    for k := range p1m {
        p1 = append(p1, k)
    }

    p2 := make([]int, 0)
    for k, v := range m {
        if v == 1 {
            p2 = append(p2, k)
        }
    }
    sort.Ints(p1)
    sort.Ints(p2)
    result := make([][]int, 0, 2)
    result = append(result, p1)
    result = append(result, p2)

    return result
}

