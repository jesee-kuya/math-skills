package skills

// StandardDeviation takes variance a float64 and returns the square root of the variance
func StandardDeviation(variance float64) float64 {
	return Sqrt(variance)
}

// Sqrt takes nb a float64 and returns the square root of the number using Newton-Raphson method.
func Sqrt(nb float64) float64 {
	s := nb / 2
	for i := 0; i <= 1000; i++ {
		s -= (s*s - nb) / (2 * s)
	}
	return s
}
