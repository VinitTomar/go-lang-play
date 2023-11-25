package main

import "fmt"


func simpleArray()  {
	var arr [2]string;
	arr[0] = "hello"
	arr[1] = "world"

	fmt.Println(arr)

	primes := [5]int {1, 2, 3, 5, 7}

	fmt.Println("array", primes)
}

func arrays()  {
	simpleArray();
}