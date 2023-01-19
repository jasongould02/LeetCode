func maxIncreaseKeepingSkyline(grid [][]int) int {
    rows, cols := getHighest(grid)
    total := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if rows[i] <= cols[j] {
                total += rows[i] - grid[i][j]
            } else if cols[j] < rows[i] {
                total += cols[j] - grid[i][j]
            }
        }
    }
    return total
}

func getHighest(grid [][]int) ([]int, []int) {
    row := make([]int, len(grid))
    col := make([]int, len(grid[0]))
    for r := 0; r < len(grid); r++ {
        for c := 0; c < len(grid[0]); c++ {
            if col[c] < grid[r][c] {
                col[c] = grid[r][c]
            }
            if row[r] < grid[r][c] {
                row[r] = grid[r][c]
            }
        }
    }
    return row, col
}
