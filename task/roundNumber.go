package task

import (
	"math"
)

func RoundNumber(num float64) float64 {
	return math.Round(num*10) / 10
}
