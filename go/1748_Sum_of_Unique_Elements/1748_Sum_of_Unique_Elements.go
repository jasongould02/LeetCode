func sumOfUnique(nums []int) int {
    m := make(map[int]int)
    for _, n := range nums {
        m[n] += 1
    }
    total := 0
    for n := range m {
        if m[n] == 1 {
            total += n
        }
    }
    return total
}
