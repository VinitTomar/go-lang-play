package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("At %v, %v", e.When, e.What);
}

func run() error {
	return MyError{
		When: time.Now(),
		What: "It didn't work.",
	}
}

func errors() {
	if err := run(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Everything is good.")
	}
}