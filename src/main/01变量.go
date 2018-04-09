package main

import (
	"fmt"
	"math"
)

func consts() {
	const a, b = "abx", 5
}

func enums() {
	const (
		cpp    = iota
		java
		python
		golang
		js
	)
	const (
		_  float64 = 1 << (iota * 10)
		KB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
}

func variable() {
	var a, b, c = 3, 4, "you"
	var s = "fuck"
	e := "aa"
	fmt.Printf("%d  %q %d %d\n", a, s, b, c)
	fmt.Printf(e)
}
func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func main() {
	//fmt.Println("Hello World")
	//variable()
	triangle()
}
