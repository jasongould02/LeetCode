func majorityElement(nums []int) int {
    m := nums[0]
    count := 0
    for i := 0; i < len(nums); i++ {
        if count == 0 {
            m = nums[i]
            count += 1
        } else if m == nums[i] {
            count += 1
        } else {
            count -= 1
        }
    }
    return m
}
