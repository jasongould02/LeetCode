func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := 0
    for _, v := range candies {
        if v >= max {
            max = v
        }
    }
    res := make([]bool, len(candies))
    for i, v := range candies {
        if v+extraCandies >= max {
            res[i] = true
        } else {
            res[i] = false
        }
    }
    return res
}
