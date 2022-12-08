func areOccurrencesEqual(s string) bool {
    m := make(map[rune]int)

    for _, c := range s {
        m[c-'a'] += 1
    }

    sample := m[rune(s[0])-'a']
    for k := range m {
        if m[k] != sample {
            return false
        }
    }
    return true
}
