package main

import (
	"fmt"
)

/*type (
	ByteSIze int64
)

const (
	a, b = 1, 2
	c, d
)

func main() {
	fmt.Println("Hello World")
	var a ByteSIze = 3
	var c = 4
	b := strconv.Itoa(c)
	fmt.Println(b)
	fmt.Println(a)
}*/
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

func main() {
	fmt.Println(KB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(ZB)
	fmt.Println(YB)
}
