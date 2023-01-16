func findPoisonedDuration(timeSeries []int, duration int) int {
    total := 0
    for i := 0; i < len(timeSeries); i++ {
        if i == len(timeSeries)-1 {
            total += duration
        } else {
            if timeSeries[i] + duration - 1 < timeSeries[i+1] {
                total += duration
            } else {
                total += timeSeries[i+1] - timeSeries[i]
            }
        }
    }
    return total
}
