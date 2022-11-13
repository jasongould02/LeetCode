func numMatchingSubseq(s string, words []string) int {
    total := 0
    passed := make(map[string]bool)
    NextWord:
    for i := 0; i < len(words); i++ {
        _, exists := passed[words[i]]
        if exists {
            if passed[words[i]] {
                total++
            }
            continue
        }

        if len(words[i]) > len(s) {
            continue
        }

        j := 0
        for k, _ := range s {
            if s[k] == words[i][j] {
                j++
            }
            if j >= len(words[i]) {
                total++
                passed[words[i]] = true
                continue NextWord
            }
        }
        passed[words[i]] = false
    }

    return total
}
