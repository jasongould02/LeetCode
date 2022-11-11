func lastStoneWeight(stones []int) int {
    sort.Ints(stones)
    var result []int
    result = stones
    for len(result) > 1 {
        x := result[len(result)-2]
        y := result[len(result)-1]
        
        if x == y {
            result = result[:len(result)-2]
        } else if x != y {
            result = result[:len(result)-1]
            result[len(result)-1] = y - x
        }
        sort.Ints(result)
    }
    if len(result) == 1 {
        return result[0]
    }
    return 0
}
