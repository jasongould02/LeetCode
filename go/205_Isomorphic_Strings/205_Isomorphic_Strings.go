func isIsomorphic(s string, t string) bool {
    length := len(s)
    var ts [256]int
    var st [256]int

    for i := range ts {
        ts[i] = -1
        st[i] = -1
    }

    for i := 0; i < length; i++ {
        if ts[t[i]] == -1 && st[s[i]] == -1 {
            ts[t[i]] = int(s[i]-'a')
            st[s[i]] = int(t[i]-'a')
        } else if ts[t[i]] != int(s[i]-'a') && st[s[i]] != int(t[i]-'a') {
            return false
        }
    }

    return true
}
