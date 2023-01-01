var seen [][]bool
var grid [][]int

func maxAreaOfIsland(grid [][]int) int {
    seen = make([][]bool, len(grid))
    for a := range seen {
        seen[a] = make([]bool, len(grid[0]))
    }
    result := 0
    for r := range grid {
        for c := range grid[r] {
            if cur := findArea(r, c, grid); cur > result {
                result = cur
            }
        }
    }
    return result
}

func findArea(row, column int, grid [][]int) int {
    if row < 0 || row >= len(grid) || column < 0 || column >= len(grid[0]) || seen[row][column] || grid[row][column] == 0 {
        return 0
    }
    seen[row][column] = true
    return 1 + findArea(row+1, column, grid) + findArea(row-1, column, grid) + findArea(row, column-1, grid) + findArea(row, column+1, grid)
}
