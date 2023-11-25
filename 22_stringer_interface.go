package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old.", p.Name, p.Age);
}

type IPAddr [4]byte;

func (ipa IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipa[0], ipa[1], ipa[2], ipa[3]);
}

func stringer() {
	p1 := Person{
		Name: "Tejas",
		Age: 5,
	}

	p2 := Person {"Garvit", 2}

	fmt.Println(p1, p2);

	hosts := map[string]IPAddr{
		"loopback": {1,1,1,1},
		"google":{172,11,43,3},
	}

	for name, ip := range hosts {
		fmt.Println(name, ":", ip)
	}
}