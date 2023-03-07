import "sort"
func sortTheStudents(score [][]int, k int) [][]int {
    sort.SliceStable(score[:], func(a, b int) bool {
        if score[a][k] - score[b][k] > 0 {
            return true
        } else {
            return false
        }
    })
    return score
}
