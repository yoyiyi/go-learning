package main

import "fmt"

func printfArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var a [5]int
	b := [3]int{1, 2, 3}
	c := [...]int{4, 5, 6, 7}
	var grid [4][5]int
	fmt.Println(grid)
	fmt.Println(a, b, c)

	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
	for i, v := range b {
		fmt.Println(i, v)
	}
}
