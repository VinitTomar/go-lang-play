package main

import "fmt"

func counting() {
	fmt.Println("Counting ...");

	for i:=0; i<10; i++ {
		defer fmt.Println(i);
	}

	fmt.Println("Done")
}

func deferMe() {
	defer fmt.Println("World!");

	fmt.Print("Hello ");

	counting();
}