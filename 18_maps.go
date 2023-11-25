package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

type Vertex2 struct {
	Lat, Long float64
}

func example1() {
	var m map[string]Vertex2 = make(map[string]Vertex2);

	m["Bell labs"] = Vertex2{
		Lat: 1234.556,
		Long: 343.22,
	}

	fmt.Println(m)
}

func mapLiteral() {
	mp := map[string]Vertex2{
		"Bell labs": {
			Lat: 4535,
			Long: 64534,
		},
		"Google": {
			Lat: 9843,
			Long: 23534,
		},
	}

	fmt.Println(mp)
}

func mapMutation() {
	mp := make(map[string]int);

	mp["key1"] = 1;
	mp["key2"] = 2;

	fmt.Println(mp);

	elm := mp["key2"];
	fmt.Println("elm2 =", elm);

	delete(mp, "key2");

	elm2, ok := mp["key2"];

	if ok {
		fmt.Println("elm2 exist =", elm2);
	} else {
		fmt.Println("elm 2 deleted")
	}

	fmt.Println(mp)

}

func countWords(s string)map[string]int {
	words := strings.Fields(s);

	counts := make(map[string]int);

	for _, word := range words {
		val, ok := counts[word];

		if ok {
			counts[word] = val + 1;
		} else {
			counts[word] = 1;
		}
	}

	return counts;
}

func maps() {
	example1();
	mapLiteral();
	mapMutation();
	wc.Test(countWords)
}