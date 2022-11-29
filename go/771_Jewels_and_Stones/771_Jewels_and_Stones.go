func numJewelsInStones(jewels string, stones string) int {
    m := make(map[rune]int)
    result := 0
    for _, l := range stones {
        m[l] += 1
    }
    for _, l := range jewels {
        amount, exists := m[l]
        if exists {
            result += amount
        }
    }
    return result
}
