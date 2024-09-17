package staticalcul

import (
	"math"
)

//Welford's online algorithm for calculating average, varince and standard deviation
func WelfordAlgo(data []int) (int, int, int) {
	var n float64
	var average float64
	var sum2 float64
	for _, d:=range data{
		n += 1
		D := float64(d)- average
		average +=  D/n
		sum2 +=  D * (float64(d) - average)
	}
	variance := sum2 / n
	stDeviation := math.Sqrt(variance)
	
	avRound := int(math.Round(average))
	vaRound := int(math.Round(variance))
	stdRound := int(math.Round(stDeviation))

	return avRound, vaRound, stdRound
}