func finalValueAfterOperations(operations []string) int {
    res := 0
    for _, v := range operations {
        res += (44 - int(v[1]))
    }
    return res
}
