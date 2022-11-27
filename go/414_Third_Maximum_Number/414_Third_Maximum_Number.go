func thirdMax(nums []int) int {
    sort.Ints(nums)
    prev := nums[len(nums)-1]
    count := 0
    for x := len(nums)-2; x >= 0; x-- {
        if nums[x] != prev {
            count += 1
            prev = nums[x]
        }
        if count == 2 {
            return prev
        }
    }
    return nums[len(nums)-1]
}
