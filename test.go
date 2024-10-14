package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	orderedAmount := decimal.NewFromFloat(123456789129991209.12345678)
	ret, err := orderedAmount.Float64()
	aa := fmt.Sprintf("%.2f", ret)
	println(err, ret, aa)
}
