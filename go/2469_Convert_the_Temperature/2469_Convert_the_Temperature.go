func convertTemperature(celsius float64) []float64 {
    answer := []float64{celsius + 273.15, (celsius*1.80) + 32.00}
    return answer
}
