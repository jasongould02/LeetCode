func heightChecker(heights []int) int {
    temp := make([]int, len(heights))
    copy(temp, heights)
    sort.Ints(temp)
    total := 0
    for i, n := range heights {
        if temp[i] != n {
            total += 1
        }
    }
    return total
}
