package main

import "fmt"

func forLoop() {
	sum := 0

	for i :=0; i <= 10; i++ {
		sum += i;
	}

	fmt.Println("Sum = ", sum)

	forLoop2()
}

func forLoop2() {
	sum := 1

	for sum < 1000 {
		sum += sum
	}

	fmt.Println("Sum2 = ", sum)
}