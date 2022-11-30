func findDuplicates(nums []int) []int {
    if len(nums) == 1 {
        return []int{}
    }
    sort.Ints(nums)
    result := make([]int, 0)
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] == nums[i+1] {
            result = append(result, nums[i])
        }
    }
    return result
}
