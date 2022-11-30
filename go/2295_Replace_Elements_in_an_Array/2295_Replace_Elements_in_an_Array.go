func arrayChange(nums []int, operations [][]int) []int {
    m := make(map[int]int, len(nums))
    for i, v := range nums {
        m[v] = i // set original position
    }
    for i := 0; i < len(operations); i++ {
        origpos := m[operations[i][0]]
        nums[origpos] = operations[i][1]
        delete(m, operations[i][0]) // remove old number & position pair
        m[operations[i][1]] = origpos
    }
    return nums
}
