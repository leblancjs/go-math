package stats

import (
	"github.com/leblancjs/log330-math-go/math"
	math2 "math"
)

func Sum(values []float64) float64 {
	sum := float64(0)
	for _, v := range values {
		sum += v
	}
	return sum
}

func Mean(values []float64) float64 {
	sum := Sum(values)
	count := float64(len(values))
	mean := sum / count
	return mean
}

func Tss(values []float64) float64 {
	mean := Mean(values)
	tss := float64(0)
	for _, v := range values {
		d := math.Distance(v, mean)
		tss += math.Square(d)
	}
	return tss
}

func Variance(values []float64) float64 {
	tss := Tss(values)
	count := float64(len(values) - 1)
	return tss / count
}

func StdDev(values []float64) float64 {
	variance := Variance(values)
	return math2.Sqrt(variance)
}
