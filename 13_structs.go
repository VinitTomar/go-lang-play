package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func structs() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2};
	v.X = 33

	fmt.Println(v)

	p := &v;
	p.Y = 1e9;

	fmt.Println(v)

	var (
		v1 = Vertex{3, 4}
		v2 = Vertex{X: 7}
		v3 = Vertex{}
		p1 = &Vertex{1, 2}
	)

	fmt.Println(v1, v2, v3, p1)

}