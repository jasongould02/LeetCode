func climbStairs(n int) int {
    var total, a, b int;
    total = 0
    a = 1
    b = 1
    if n == 1 {
        return 1
    }

    for i := 2; i <= n; i++ {
        total = a + b
        a = b
        b = total
    }
    return total
}

