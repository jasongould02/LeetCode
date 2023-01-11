func findRepeatedDnaSequences(s string) []string {
    var results []string
    m := make(map[string]int)

    for i := 0; i < len(s)-9; i++ {
        str := s[i:i+10]
        if m[str] == 1 {
            results = append(results, str)
        }
        m[str]++
    }

    return results
}
