func simplifyPath(path string) string {
    res := ""
    split := strings.Split(path, "/")

    for _, v := range split {
        if v == "" || v == "." {
            continue
        }
        if v != ".." {
            res = res + "/" + v
        } else if len(res) > 0 {
            res = res[:strings.LastIndex(res, "/")]
        }
    }

    if strings.HasPrefix(res, "/") {
        return res
    } else {
        return "/" + res
    }
}
