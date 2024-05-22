package skills

// Median sorts the slice numbers from low to highest and returns the value at the middle
func Median(numbers []float64) float64 {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	if len(numbers)%2 != 0 {
		n := len(numbers) - 1
		return numbers[n/2]
	} else {
		n1 := len(numbers) / 2
		n2 := n1 + 1
		return (numbers[n1-1] + numbers[n2-1]) / 2
	}
}
