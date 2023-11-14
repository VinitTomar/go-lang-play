package main

import "fmt"

func add(x int, y int) int {
	return x + y;
}

func add2(x, y int) int {
	return x + y;
}

func swap(a, b string) (string, string) {
	return b, a
}

func namedReturnSplit(sum int) (x, y int) {
	x = sum * 4/9
	y = sum - x
	return
}

// var c, java, python bool;
var c, java, python = 1, false, "!no";

func showAdd() {
	fmt.Println("Sum of 2 and 3 is: ", add(2,3));
	fmt.Println("Sum of 5 and 18 is: ", add2(5, 18));

	var a, b string = swap("Hello", "world");
	fmt.Println("Swap: ", a, b);

	num := 100
	x, y := namedReturnSplit(num);
	fmt.Println("Split", num, " => ", x, y)

	var i int;
	fmt.Println("Variables", i, c, java, python)
}