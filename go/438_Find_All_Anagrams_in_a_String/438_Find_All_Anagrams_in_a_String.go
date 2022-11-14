func findAnagrams(s string, p string) []int {
    var pset [26]int
    var subset [26]int
    for _, v := range p {
        pset[v-'a'] += 1
    }
    result := make([]int, 0)
    for i := 0; i < len(s); i++ {
        subset[s[i]-'a']++
        if i >= len(p) {
            subset[s[i-len(p)]-'a']--
        }
        if subset == pset {
            result = append(result, i-len(p)+1)
        }
    }
    return result
}
