func checkAlmostEquivalent(word1 string, word2 string) bool {
    letters := make([]int, 26)
    for _, c := range word1 {
        letters[c-'a'] += 1
    }

    for _, c := range word2 {
        letters[c-'a'] -= 1
    }

    for _, c := range letters {
        if abs(c) > 3 {
            return false
        }
    }
    return true
}

func abs(a int) int {
    if a < 0 {
        return -1 * a
    } else {
        return a
    }
}
