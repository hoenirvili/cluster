package util

import "math"

// Round rounds the floating point
// number based on the prec given
func Round(x float64, prec int) float64 {
	if x == 0 {
		return 0
	}
	if prec >= 0 && x == math.Trunc(x) {
		return x
	}

	pow := math.Pow10(prec)
	intermed := x * pow
	if math.IsInf(intermed, 0) {
		return x
	}
	if x < 0 {
		x = math.Ceil(intermed - 0.5)
	} else {
		x = math.Floor(intermed + 0.5)
	}

	if x == 0 {
		return 0
	}

	return x / pow
}
