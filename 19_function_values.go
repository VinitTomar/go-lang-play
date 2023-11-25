package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4);
}

func adder() func(int) int {
	sum := 0;

	return func(i int) int {
		sum += i;
		return sum;
	}
}

func closureExample() {
	fmt.Println("Closure example");

	pos, neg := adder(), adder()

	for i := 0; i<10; i++ {
		fmt.Printf("pos %d, neg %d\n", pos(i), neg(-2*i))
	}
}

func fibonacci() func()int {
	a, b := 0, 1;

	computeValues := func() {
		sum := a + b;
		a = b;
		b = sum;
	}

	return func() int {
		if a == 0 {
			computeValues();
			return 0;
		}
	
		computeValues();

		return a;
	}
}

func functionValues() {
	hypot := func(a, b float64) float64 {
		return math.Sqrt(a*a + b*b);
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	closureExample();

	fmt.Println("Fibonacci up to 10 places");
	fib := fibonacci();

	for i :=0; i<10; i++ {
		fmt.Print(fib(), " ")
	}
	fmt.Println()
}