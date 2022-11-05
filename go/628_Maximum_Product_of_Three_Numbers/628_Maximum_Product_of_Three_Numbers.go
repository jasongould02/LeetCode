func maximumProduct(nums []int) int {
    sort.Ints(nums)
    top := nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
    bottom := nums[0] * nums[1] * nums[len(nums)-1]

    if top > bottom {
        return top
    } else {
        return bottom
    }
}
