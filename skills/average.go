package skills

// Mean takes a slice numbers of type float64 and returns the average of the numbers as float64
func Mean(numbers []float64) float64 {
	var num float64
	var count float64
	for i := 0; i < len(numbers); i++ {
		num += numbers[i]
		count++
	}
	return num / count
}
