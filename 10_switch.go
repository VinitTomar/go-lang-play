package main

import (
	"fmt"
	"runtime"
	"time"
)

func printOS() {
	fmt.Print("Go runs on ");

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X");
	case "linux":
		fmt.Println("Linux");
	default:
		fmt.Printf("%s \n", os);
	}
}

func whenIsSaturday() {
	fmt.Print("When is Saturday?");

	today := time.Now().Weekday();

	switch time.Saturday {
	case today + 0:
		fmt.Println("It is today.");
	case today + 1:
		fmt.Println("It is tomorrow.");
	case today + 2:
		fmt.Println("It is day after tomorrow");
	default:
		fmt.Println("It is too far away.");
	}
}

func sayWhat() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!");
	case t.Hour() < 17:
		fmt.Println("Good afternoon.");
	default:
		fmt.Println("Good evening.")

	}
}

func switchStatement() {
	printOS();
	whenIsSaturday();
	sayWhat();
}