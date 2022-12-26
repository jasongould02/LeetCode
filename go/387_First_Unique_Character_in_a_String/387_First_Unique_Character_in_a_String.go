func firstUniqChar(s string) int {
    m := make(map[int]int)
    for _, v := range s {
        m[int(v - 'a')] += 1
    }
    index := -1
    for _, v := range s {
        if m[int(v-'a')] == 1 {
            index = strings.Index(s, string(v))
            break
        }
    }
    return index
}
