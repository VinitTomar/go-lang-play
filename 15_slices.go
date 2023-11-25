package main

import (
	"fmt"
	"strings"
)

func simpleSlice() {
	primes := [5]int {1, 2, 3, 5, 7}
	var slice []int = primes[1:4]
	fmt.Println("slice", slice)
	slice2 := primes[2:5]
	slice[1] = 44
	fmt.Println("Slice2 ", slice2);
}

func sliceLiteral() {
	s1 := []int {1,2,3,4,5};
	s2 := []struct{
		age int 
		name string 
	}{
		{age: 22, name: "name_1"},
		{age: 34, name: "name_2"},
	}

	fmt.Println("slice literal",s1,"\nslice literal struct", s2)
}

func sliceMake() {
	fmt.Println("Make few slices");
	a := make([]int, 5);
	printSlice("a", a);

	b := make([]int, 0, 5);
	printSlice("b", b);

	c := b[:2]
	printSlice("c", c);

	d := c[2:5]
	printSlice("d", d);

}

func printSlice(name string, s []int) {
	fmt.Printf("%s len=%d, cap=%d, %v\n", name, len(s), cap(s), s)
}

func makeTicTacToeBoard() {
	fmt.Println("Tic Tac Toe board");

	board := [][]string {
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i:=0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "));
	}
}

func appendSlice() {
	fmt.Println("Append slice");

	var s []int;
	printSlice("Append", s);

	s = append(s, 0);
	printSlice("Append", s);

	s = append(s, 1, 2, 3)
	printSlice("Append", s)

}

func slices() {
	simpleSlice();
	sliceLiteral();
	sliceMake();
	makeTicTacToeBoard();
	appendSlice();
}