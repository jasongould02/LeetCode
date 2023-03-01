func majorityElement(nums []int) []int {
    l := len(nums)
    m := make(map[int]int, l)
    res := make([]int, 0)
    NEXT:
    for _, v := range nums {
        m[v] += 1
        if m[v] > (l/3) {
            for _, k := range res {
                if k == v {
                    continue NEXT
                }
            }
            res = append(res, v)
        }
    }
    return res
}
