func uniquePaths(m int, n int) int {
    grid := make([][]int, m)
    for i := range grid {
        grid[i] = make([]int, n)
    }
    return getPath(m-1, n-1, grid)
}

func getPath(m, n int, prev_grid [][]int) int {
    if m < 0 || n < 0 {
        return 0
    } else if m == 0 || n == 0 {
        return 1
    } else if prev_grid[m][n] > 0 {
        return prev_grid[m][n]
    } else {
        prev_grid[m][n] = getPath(m-1, n, prev_grid) + getPath(m, n-1, prev_grid)
        return prev_grid[m][n]
    }
}
