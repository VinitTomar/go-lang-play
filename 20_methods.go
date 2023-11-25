package main

import (
	"fmt"
	"math"
)

type Vertex3 struct {
	X float64
	Y float64
}

func (v Vertex3) Abs() float64 {
	x := v.X;
	y := v.Y;
	return math.Sqrt(x*x + y*y);
}

type MyFloat float64;

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f);
	}

	return float64(f);
}

func (v *Vertex3) scale(f float64) {
	v.X = v.X * f;
	v.Y = v.Y * f;
}

func methods() {
	v := Vertex3 {
		X: 3,
		Y: 4,
	}

	fmt.Println(v.Abs())

	mf := MyFloat(-math.Sqrt2)
	fmt.Println(mf.Abs())

	v.scale(2);

	fmt.Println(v);
}