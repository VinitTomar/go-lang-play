package main

import "fmt"


func pointers() {
	i, j := 42, 2700
	p := &i

	fmt.Println("i = ", *p);

	*p = 21;

	fmt.Println("i = ", i);
	fmt.Println("j = ", j);

	p = &j;
	*p = *p / 10;

	fmt.Println("j = ", j);
}