package main

import (
	"fmt"
	"time"
)

func fibonacciSelect(c, quit chan int) {

	x, y := 0, 1;

	for {
		select {
		case c <- x:
			x, y = y, x+y;
		
		case <- quit:
			fmt.Println("quit select fibonacci");
			return;
		}
	}

}

func defaultSelect() {
	tick := time.Tick(100 * time.Millisecond);
	boom := time.After(500 * time.Millisecond);

	for {
		select {
		case <- tick:
			fmt.Println("tick.");
		case <- boom:
			fmt.Println("boom!!!!");
			return;
		default:
			fmt.Println("     .");
			time.Sleep(50 * time.Millisecond);
		}
	}

}

func selects() {
	c := make(chan int);
	quit := make(chan int);

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Print(<-c, " ,")
		}
		fmt.Println();
		quit <- 1;
	}()

	fibonacciSelect(c, quit);
	defaultSelect();
}