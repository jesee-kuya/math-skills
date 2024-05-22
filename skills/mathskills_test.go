package skills

import (
	"testing"
)

var testcases = []struct {
	name              string
	numbers           []float64
	median            float64
	average           float64
	variance          float64
	standardDeviation float64
}{
	{"Test1", []float64{5, 7, 3, 7}, 6, 5.5, 3.6666666666666665, 1.9148542155126762},
}

func TestMathSkills(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// test for median
			expectedMedian := tc.median
			gotMedian := Median(tc.numbers)
			// test for average
			expectedAverage := tc.average
			gotAverage := Mean(tc.numbers)
			// test for variance
			expectedVariance := tc.variance
			gotVariance := Variance(tc.numbers, gotAverage)
			// test for standard deviation
			expectedStandardDeviation := tc.standardDeviation
			gotStandardDeviation := StandardDeviation(gotVariance)

			switch {
			case expectedMedian != gotMedian:
				t.Errorf("Expected %v but got %v\n", expectedMedian, gotMedian)
			case expectedAverage != gotAverage:
				t.Errorf("Expected %v but got %v\n", expectedAverage, gotAverage)
			case expectedVariance != gotVariance:
				t.Errorf("Expected %v but got %v\n", expectedVariance, gotVariance)
			case expectedStandardDeviation != gotStandardDeviation:
				t.Errorf("Expected %v but got %v\n", expectedStandardDeviation, gotStandardDeviation)
			}
		})
	}
}
