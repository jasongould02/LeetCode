func addDigits(num int) int {
    for num > 9 {
        num = sumDigits(num)
    }
    return num
}

func sumDigits(n int) int {
    total := 0
    for n > 0 {
        digit := n % 10
        n = n / 10
        total += digit
    }
    return total
}
