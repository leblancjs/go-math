package math

func Distance(a float64, b float64) float64 {
	return b - a
}

func Power(v float64, p uint) float64 {
	if p == 0 {
		return 1
	}

	return v * Power(v, p-1)
}

func Square(v float64) float64 {
	return Power(v, 2)
}
