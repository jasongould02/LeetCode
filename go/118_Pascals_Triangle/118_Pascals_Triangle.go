func generate(numRows int) [][]int {
    tri := make([][]int,0)
    if numRows == 0 {
        return tri
    }
    i := 2
    tri = append(tri, []int{1})
    if numRows == 1 {
        return tri
    }
    tri = append(tri, []int{1,1})
    if numRows == 2 {
        return tri
    }
    for i < numRows {
        newRow := make([]int, i+1)
        newRow[0], newRow[len(newRow)-1] = 1, 1
        for j := 1; j < len(newRow)-1; j++ {
            newRow[j] = tri[i-1][j] + tri[i-1][j-1]
        }
        tri = append(tri, newRow)
        i++
    }
    return tri
}
