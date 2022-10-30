func countPoints(points [][]int, queries [][]int) []int {
    answer := make([]int, len(queries))
    for i, v := range queries {
        for _, k := range points {
            if ((k[0]-v[0])*(k[0]-v[0]))+((k[1]-v[1])*(k[1]-v[1])) <= v[2]*v[2] {
                answer[i]++
            }
        }
    }
    return answer
}
