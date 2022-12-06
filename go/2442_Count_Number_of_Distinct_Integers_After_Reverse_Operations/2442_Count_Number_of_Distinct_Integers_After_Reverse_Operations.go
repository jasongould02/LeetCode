func countDistinctIntegers(nums []int) int {
    m := make(map[int]int)
    result := 0
    for _, v := range nums {
        _, exists := m[v]
        if !exists {
            m[v] = 1
            result += 1
        }
        reverse := reverseNumber(v)
        _, exists = m[reverse]
        if !exists {
            m[reverse] = 1
            result += 1
        }
    }
    return result
}

func reverseNumber(number int) int {
    reversed := 0
    for number > 0 {
        reversed = reversed*10 + number%10
        number /= 10
    }
    return reversed
}
