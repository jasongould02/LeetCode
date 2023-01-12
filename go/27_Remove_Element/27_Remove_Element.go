func removeElement(nums []int, val int) int {
    insertIndex := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[insertIndex] = nums[i]
            insertIndex++
        }
    }
    return insertIndex
}
