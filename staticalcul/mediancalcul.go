package staticalcul

import (
	"math"
	"sort"
)

func MedianCalcul(data []int) int {
	sort.Ints(data)
	n := len(data)
	var res float64
	if  n % 2 == 1{
		res = float64(data[n/2])
	}else{
		res = float64(data[n/2-1]+data[n/2]) / 2
	}
	return int(math.Round(res))
}
