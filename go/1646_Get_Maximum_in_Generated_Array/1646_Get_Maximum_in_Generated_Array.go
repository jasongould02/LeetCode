func getMaximumGenerated(n int) int {
    if n == 2 || n == 1 {
        return 1
    } else if n == 0 {
        return 0
    }
    nums := make([]int, n+1)
    nums[1] = 1
    max := 0
    for i := 1; i < n+1; i++ {
        if i*2 >= 2 && i*2 < n {
            nums[i*2] = nums[i]
            if nums[i*2] > max {
                max = nums[i*2]
            }
        }
        if ((i*2)+1) >= 2 && ((i*2)+1) <= n {
            nums[(2*i)+1] = nums[i] + nums[i+1]
            if nums[(2*i)+1] > max {
                max = nums[(2*i)+1]
            }
        }
    }
    return max
}
