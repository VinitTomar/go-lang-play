package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64;
}

type MyFloat2 float64;

func (f MyFloat2) Abs() float64 {
	if f < 0 {
		return float64(-f);
	}

	return float64(f);
}

type Vertex4 struct {
	X float64
	Y float64
}

func (v *Vertex4) Abs() float64 {
	if v == nil {
		fmt.Println("<nil>");
		return 0;
	}

	x := v.X;
	y := v.Y;

	return math.Sqrt(x*x + y*y);
}

func describe(i interface{}) {
	fmt.Printf("%v, %T\n", i ,i);
}

func emptyInterface() {
	var i interface{};
	describe(i);

	i = 3;
	describe(i);

	i = "hello"
	describe(i);
}

func typeAssertion() {
	var i interface{} = "hello";

	vInt, ok := i.(int);

	if ok {
		fmt.Println("Value is Int =", vInt);
	}

	vStr, oks := i.(string);

	if oks {
		fmt.Println("Value is string =", vStr);
	}

	// f := i.(float64) // panic
	// fmt.Println(f)
}

func switchesTypes() {
	var i interface{} = 12.3;

	switch v := i.(type) {
	case int:
		fmt.Printf("Twice of %d is %d\n", v, v*2);
	case string:
		fmt.Printf("Length of string %s is %d\n", v, len(v));
	default:
		fmt.Printf("Unsupported type %T\n", i)
	}
}

func interfaces() {
	fmt.Println("Interfaces examples.");

	var a Abser;

	f := MyFloat2(-math.Sqrt2);
	v := Vertex4 {
		X: 3,
		Y: 4,
	}

	a = f;
	fmt.Println(a.Abs());
	
	a = &v;
	fmt.Println(a.Abs());

	var v2 *Vertex4;
	a = v2;
	fmt.Println(a.Abs());

	emptyInterface();
	typeAssertion();
	switchesTypes();
}