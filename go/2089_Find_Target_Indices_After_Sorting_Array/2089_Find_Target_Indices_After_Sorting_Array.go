import "sort"
func targetIndices(nums []int, target int) []int {
    sort.Ints(nums)
    out := make([]int, 0, 10)
    left := 0
    right := len(nums) - 1

    for left <= right {
        mid := (left + (right - left) / 2)
        if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid
            break
        }
    }
    if left < len(nums) && nums[left] == target {
        for i, v := range nums[:right+1] {
            if v == target {
                out = append(out, i)
            }
        }
    } else {
        return []int{}
    }
    return out
}
