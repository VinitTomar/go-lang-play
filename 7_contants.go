package main

import "fmt"

const world = "World";

const (
	big = 1 << 100
	small = big >> 99
)

func needInt(num int) int {
	return num * 10 + 1
}

func needFloat(num float64) float64 {
	return num * 0.1
}

func constants() {
	const PI = 3.14;

	fmt.Println("Hello ", world);
	fmt.Println("PI ", PI);

	fmt.Println("Small int", needInt(small));
	fmt.Println("Small float ", needFloat(small));

	// fmt.Println("Big int", needInt(big)); // overflow
	fmt.Println("Big float", needFloat(big));

}