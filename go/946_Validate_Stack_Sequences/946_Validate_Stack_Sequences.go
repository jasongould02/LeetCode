func validateStackSequences(pushed []int, popped []int) bool {
    stack := make([]int, 0, 0)
    length := len(pushed)
    pos1, pos2 := 0, 0
    for pos1 < length {
        stack = append(stack, pushed[pos1])
        pos1 += 1
        for len(stack) > 0 && pos2 < length && popped[pos2] == stack[len(stack) - 1] {
            stack = stack[:len(stack)-1]
            pos2 += 1
        }
    }
    return len(stack) == 0
}
