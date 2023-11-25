package main

import (
	"fmt"
	"math"
)


type ErrNegativeSqrt float64;

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Negative numbers are not supported %v", float64(e));
}

func sqrt2(n float64) (float64, error) {
	if(n < 0) {
		return 0, ErrNegativeSqrt(n)
	}

	return math.Sqrt(n), nil;
}

func errorsSqrt() {
	fmt.Println(sqrt2(2));
	fmt.Println(sqrt2(-2));
}