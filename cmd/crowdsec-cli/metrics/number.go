package metrics

import (
	"fmt"
	"math"
)

type unit struct {
	value  int64
	symbol string
}

var ranges = []unit{
	{value: 1e18, symbol: "E"},
	{value: 1e15, symbol: "P"},
	{value: 1e12, symbol: "T"},
	{value: 1e9, symbol: "G"},
	{value: 1e6, symbol: "M"},
	{value: 1e3, symbol: "k"},
	{value: 1, symbol: ""},
}

func formatNumber(num int) string {
	goodUnit := unit{}

	for _, u := range ranges {
		if int64(num) >= u.value {
			goodUnit = u
			break
		}
	}

	if goodUnit.value == 1 {
		return fmt.Sprintf("%d%s", num, goodUnit.symbol)
	}

	res := math.Round(float64(num)/float64(goodUnit.value)*100) / 100

	return fmt.Sprintf("%.2f%s", res, goodUnit.symbol)
}
