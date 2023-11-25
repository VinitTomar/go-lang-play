package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (mr MyReader) Read(b []byte) (int, error) {
	ln := len(b);
	cp := cap(b);
	
	fmt.Println("length", ln, "capacity", cp);
	
	for i:=0; i<ln; i++ {
		b[i] = 'A';
	}
	
	return ln, nil;
}


func readerExr1() {
	reader.Validate(MyReader{})
}
