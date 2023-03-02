func mostFrequentEven(nums []int) int {
    m := make(map[int]int, 0)
    curNumber := -1
    curNumberFreq := 0
    for _, v := range nums {
        if (v & 1) == 0 {
            m[v] += 1

            if m[v] == curNumberFreq {
                if v < curNumber {
                    curNumber = v
                    curNumberFreq = m[v]
                    continue
                }
            }

            if m[v] > curNumberFreq {
                curNumber = v
                curNumberFreq = m[v]
            }
        }
    }
    return curNumber
}
