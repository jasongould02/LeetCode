func fib(n int) int {
    list := make([]int, 2)
    list[0] = 0
    list[1] = 1
    
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    }
    
    for i := 2; i < n+1; i++ {
        list = append(list, list[i-1] + list[i-2])
    }
    
    return list[len(list)-1]    
}