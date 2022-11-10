func numIslands(grid [][]byte) int {
    count := 0
    for x := 0; x < len(grid); x++ {
        for y := 0; y < len(grid[0]); y++ {
            if grid[x][y] == '1' {
                removeIsland(grid, x, y)
                count += 1
            }
        }
    }
    return count
}

func removeIsland(grid [][]byte, x, y int) {
    if grid[x][y] != '1' {
        return
    }
    grid[x][y] = '0';
    if x+1 < len(grid) {
        removeIsland(grid, x+1, y)
    }
    if x-1 >= 0 {
        removeIsland(grid, x-1, y)
    }
    if y+1 < len(grid[0]) {
        removeIsland(grid, x, y+1)
    }
    if y-1 >= 0 {
        removeIsland(grid, x, y-1)
    }
}
