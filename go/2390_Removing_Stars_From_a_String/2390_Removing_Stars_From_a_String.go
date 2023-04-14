func removeStars(s string) string {
    stack := make([]byte, 0, len(s))

    for i, _ := range s {
        if s[i] == '*' {
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, s[i])
        }
    }

    return string(stack)
}
