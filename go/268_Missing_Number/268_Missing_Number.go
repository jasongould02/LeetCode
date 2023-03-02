import "sort"

func missingNumber(nums []int) int {
    sort.Ints(nums)
    j := 0
    for i, v := range nums {
        if v != j {
            return i
        }
        j++
    }
    return j
}
