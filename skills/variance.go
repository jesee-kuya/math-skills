package skills

// Variance takes numbers a slice and mean of float64 
// Variance returns the variance of the numbers in float64
func Variance(numbers []float64, mean float64) float64 {
	n := float64(len(numbers))
	var num float64
	for i := 0; i < len(numbers); i++ {
		numbers[i] -= mean
		numbers[i] *= numbers[i]
		num += numbers[i]
	}
	return num / n 
}
