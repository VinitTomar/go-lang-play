package main

import "fmt"

func Index[T comparable](sl []T, num T) int {
	for i, v := range sl {
		if v == num {
			return i;
		}
	}

	return -1;
}

type List[T any] struct {
	next *List[T]
	val T
}

type IterateList interface {
	iterate();
}

func (l List[T])iterate() {
	itm := &l;

	for {
		fmt.Println("List item is ", itm.val);
		itm = itm.next;

		if itm == nil {
			break;
		}

	}
}

func iterateList() {
	lst := &List[int] {
		val: 1,
		next: &List[int]{
			val: 2,
			next: &List[int]{
				val: 3,
				next: nil,
			},
		},
	};

	var lstItr IterateList = lst;

	lstItr.iterate();
}

func genericFunctionParams() {
	intArr := []int{1, 3, 5, 33, 65}
	fmt.Println("intArr: Index of '5' is ", Index(intArr, 5));

	strArr := []string{"aaa", "bbb", "ccc", "ddd", "eee"}
	fmt.Println("strArr: Index of 'bbb' is ", Index(strArr, "bbb"));
}

func generics() {
	genericFunctionParams();
	iterateList();
}