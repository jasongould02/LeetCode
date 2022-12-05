func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int, len(nums))
    for _, n := range nums {
        m[n] += 1
    }

    keyList := make([]int, len(m))

    i := 0
    for k := range m {
        keyList[i] = k
        i++
    }

    sort.Slice(keyList, func(i int, j int) bool {
        return m[keyList[i]] > m[keyList[j]]
    })

    return keyList[:k]
}
