package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)


func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch);
	}
	
	ch <- t.Value;
	
	if t.Right != nil {
		Walk(t.Right, ch);
	}
	
}


func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int);
	go func() {
		Walk(t1, ch1);
		close(ch1);
	}();
	
	ch2 := make(chan int);
	go func() {
		Walk(t2, ch2);
		close(ch2);
	}();
		
	for {
		i, iok := <-ch1;

		j, jok := <-ch2;

		if !iok || !jok {
			break;
		}
		
		if i != j {
			return false;
		}
	}
	
	return true;
}

func binaryTreeEql() {
	fmt.Println(Same(tree.New(1), tree.New(1)))

}