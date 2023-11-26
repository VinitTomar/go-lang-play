package main

import "fmt"

func sum(sl []int, ch chan int) {
	sum := 0;

	for _, v := range sl {
		sum += v;
	}

	fmt.Printf("Sum of %v is = %d\n", sl, sum);

	ch <- sum;
}

func simpleChannel() {
	sl := []int { 1, 2, 3, 4, 5, -1, -2, -3, -4, -5};
	
	c := make(chan int);
	
	sl1 := sl[:len(sl)/2];
	sl2 := sl[len(sl)/2:]
	
	fmt.Println("first slice = ", sl1);
	fmt.Println("second slice = ", sl2);
	
	go sum(sl2, c);
	go sum(sl1, c);
	
	x, y := <-c, <-c;
	
	fmt.Printf("X=%d, Y=%d, X+Y=%d\n", x, y, x+y);
}

func bufferedChannel() {
	fmt.Println("Channel with buffer cap 2");

	ch := make(chan int, 2);

	ch <- 1;
	ch <- 2;
	
	fmt.Println(<-ch);
	ch <- 3;

	fmt.Println(<-ch);
	fmt.Println(<-ch);
}

func fibonacciChannel(n int, c chan int) {
	fmt.Println("fibonacci series channel")
	x, y := 0, 1;

	for i := 0; i < n; i++ {
		c <- x;
		x, y = y, x+y;
	}

	close(c);
}

func channels() {
	simpleChannel();
	bufferedChannel();

	ch := make(chan int, 10);

	go fibonacciChannel(cap(ch), ch);

	for i := range ch {
		fmt.Print(i, " ,");
	}

	fmt.Println();
}