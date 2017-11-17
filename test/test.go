package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	var b1, b2, b3, bigSum big.Float
	b1.SetFloat64(65.233)
	b2.SetFloat64(987.233)
	b3.SetFloat64(22.33)

	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Printf("BigSum = %.10g\n", &bigSum)

	circleRadius := 15.5
	circumference := circleRadius * math.Pi
	fmt.Printf("Circumference: %.1f\n", circumference)
}
