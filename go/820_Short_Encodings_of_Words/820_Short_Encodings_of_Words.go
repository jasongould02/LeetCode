func minimumLengthEncoding(words []string) int {
    var sb strings.Builder

    sort.Slice(words, func(i, j int) bool {
        return len(words[i]) > len(words[j])
    })

    for _, v := range words {
        if strings.Contains(sb.String(), string(v + "#")) {
            continue
        } else {
            sb.WriteString(v + "#")
        }
    }
    return len(sb.String())
}
