package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i";
	}

	return fmt.Sprint(math.Sqrt(x))
}

func powLimit(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g therefore ", v, lim)
	}

	return lim
}

func ifElseStatement() {
	fmt.Println("Sqrt 4 = ", sqrt(4));
	fmt.Println("Sqrt -4 = ", sqrt(-4));

	fmt.Println("if lim = 3000, then 2^10 = ", powLimit(2, 10, 3000))
	fmt.Println("if lim = 1000, then 2^10 = ", powLimit(2, 10, 1000))
}