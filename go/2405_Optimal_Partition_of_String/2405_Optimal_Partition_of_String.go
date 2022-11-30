func partitionString(s string) int {
    count := 1
    var m [26]int
    for _, c := range s {
        if m[c - 'a'] == 1 {
            count += 1
            m = [26]int{}
        }
        m[c - 'a'] += 1
    }
    return count
}
