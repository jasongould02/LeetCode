func diagonalSum(mat [][]int) int {
    b := len(mat)
    h := len(mat)
    primary := 0
    secondary := 0
    for i := 0; i < b; i++ {
            primary += mat[i][i]
            secondary += mat[b-i-1][i]
    }
    res := primary + secondary
    if h & 1 == 1 {
        return res - mat[b/2][h/2]
    }
    return res
}
