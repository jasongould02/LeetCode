func findTheDifference(s string, t string) byte {
    result := make([]int, 26)
    for i, v := range s {
        result[int(v-'a')] += 1
        result[int(t[i]-'a')] -= 1
    }
    result[int(t[len(t)-1]-'a')] -= 1
    out := 0
    for i, v := range result {
        if v == -1 {
            out = i
        }
    }
    return byte(out+'a')
}
