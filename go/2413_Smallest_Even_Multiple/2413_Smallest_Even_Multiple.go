func smallestEvenMultiple(n int) int {
    if n & 1 == 0 {
        return n
    } else {
        return n * 2
    }
}
