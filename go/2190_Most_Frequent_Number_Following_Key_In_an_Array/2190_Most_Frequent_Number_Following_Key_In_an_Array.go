func mostFrequent(nums []int, key int) int {
    suffix := make(map[int]int)
    result := 0
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] == key {
            suffix[nums[i+1]] += 1
            if suffix[nums[i+1]] > suffix[result] {
                result = nums[i+1]
            }
        } 
    }
    
    return result
}