func isAnagram(s string, t string) bool {
    ss := make([]int, 26)
    if len(s) != len(t) {
        return false
    }
    for i, l := range s {
        ss[l-'a'] += 1
        ss[t[i]-'a'] -= 1
    }
    for _, l := range ss {
        if l != 0 {
            return false
        }
    }
    return true
}
