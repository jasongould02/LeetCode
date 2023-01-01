func maximumDetonation(bombs [][]int) int {
    max := 0
    for i, _ := range bombs {
        cur := search(i, make([]bool, len(bombs)), bombs)
        if cur > max {
            max = cur
        }
    }
    return max
}

func search(cur int, passed []bool, bombs [][]int) int {
    count := 0
    passed[cur] = true

    for i := 0; i < len(bombs); i++ {
        if !passed[i] && inRange(bombs[cur], bombs[i]) {
            passed[i] = true
            count += search(i, passed, bombs)
        }
    }
    return count + 1
}


func inRange(b1 []int, b2 []int) bool {
    xd := (b1[0] - b2[0])
    yd := (b1[1] - b2[1])
    rd := (b1[2])

    if (xd * xd) + (yd * yd) <= rd*rd {
        return true
    }
    return false
}
