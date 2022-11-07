func arraySign(nums []int) int {
    negCount := 0
    
    for _, v := range nums {
        if v < 0 {
            negCount++
        } else if v == 0 {
            return 0
        }
    }
    
    if negCount % 2 == 0 {
        return 1
    } else {
        return -1
    }
    
}