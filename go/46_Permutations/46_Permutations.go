var result [][]int
var length int
func permute(nums []int) [][]int {
    result = make([][]int, 0)
    length = len(nums)
    gen(length, nums)
    return result
}

func gen(k int, nums []int) {
    if k == 1 {
        out := make([]int, length)
        copy(out, nums)
        result = append(result, out)
        return
    }
    gen(k - 1, nums)
    for i := 0; i < k - 1; i++ {
        if i < k-1 {
            if k & 1 == 0 {
                nums[i], nums[k-1] = nums[k-1], nums[i]
            } else {
                nums[0], nums[k-1] = nums[k-1], nums[0]
            }
        }
        gen(k - 1, nums)
    }
}
