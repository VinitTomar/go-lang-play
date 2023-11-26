package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i :=4; i<5; i++ {
		time.Sleep(100 * time.Millisecond);
		fmt.Println(s, i);
	}
}


func routines() {
	go say("world");
	go say("me");
	say("hello");

}