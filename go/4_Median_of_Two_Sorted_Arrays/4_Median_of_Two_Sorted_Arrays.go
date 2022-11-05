func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    nums1 = append(nums1, nums2...)
    sort.Ints(nums1)
    mid := len(nums1) / 2
    if len(nums1) % 2 == 1 {
        return float64(nums1[mid])
    } else {
        return float64(nums1[mid-1] + nums1[mid]) / float64(2)
    }
}
