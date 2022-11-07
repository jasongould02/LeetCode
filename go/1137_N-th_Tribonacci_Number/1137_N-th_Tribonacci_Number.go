func tribonacci(n int) int {
    list := make([]int, 3)
    list[0] = 0
    list[1] = 1
    list[2] = 1
    
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    } else if n == 2 {
        return 1
    }
    
    for i := 3; i < n+1; i++ {
        list = append(list, list[i-1] + list[i-2] + list[i-3])
    }
    
    return list[len(list)-1]  
    
}
