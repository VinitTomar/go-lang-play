package main

import "fmt"

func powers() {
	var pows []int = []int{ 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024};

	for i, v := range pows {
		fmt.Printf("2^%d = %d\n",i, v);
	}
}

func rangeExample() {
	powers();
}