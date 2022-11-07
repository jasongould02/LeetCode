func smallerNumbersThanCurrent(nums []int) []int {
    result := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums); j++ {
            if (i != j) {
                if nums[i] > nums[j] {
                    result[i] += 1
                }
            }
        }
    }
    return result
}