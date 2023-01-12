func containsNearbyDuplicate(nums []int, k int) bool {
    m := make(map[int]int, k)

    for i, vi := range nums {
        j, exists := m[vi]
        if exists && abs(i-j) <= k {
            return true
        } else {
            m[vi] = i
        }
    }

    return false
}

func abs(i int) int {
    if i < 0 {
        return i * -1
    } else {
        return i
    }
}
