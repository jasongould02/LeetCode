func rearrangeBarcodes(barcodes []int) []int {
    m := make(map[int]int)
    result := make([]int, len(barcodes))
    conv := make([][]int, 0, len(m))

    for _, n := range barcodes {
        m[n] += 1
    }

    for k, v := range m {
        conv = append(conv, []int{k, v})
    }

    sort.Slice(conv, func(i int, j int) bool {
        return conv[i][1] > conv[j][1]
    })

    cur := 0
    for i := 0; i < len(barcodes); i += 2 {
        if conv[cur][1] == 0 {
            cur += 1 // take next highest value since current highest value is now 0
        }
        result[i] = conv[cur][0]
        conv[cur][1] -= 1
    }

    for i := 1; i < len(barcodes); i += 2 {
        if conv[cur][1] == 0 {
            cur += 1 // take next highest value since current highest value is now 0
        }
        result[i] = conv[cur][0]
        conv[cur][1] -= 1
    }
    return result
}
