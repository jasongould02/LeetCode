func isHappy(n int) bool {
    m := make(map[int]bool)
    for n > 1 {
        if _, exists := m[n]; exists {
            return false
        }
        m[n] = true
        n = sumNum(n)
    }
    return true
}

func sumNum(n int) int {
    total := 0
    for n > 0 {
        digit := n % 10
        n = n / 10
        total += digit * digit
    }
    return total
}
