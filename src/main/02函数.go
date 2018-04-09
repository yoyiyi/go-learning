package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func div(a, b int) (q, r int) {
	return a / b, a % b
}
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println("Calling func %s with args(%d,%d).", opName, a, b)
	return op(a, b)
}

func sum(nums ...int) int {
	s := 0
	for i := range nums {
		s += nums[i]
	}
	return s
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

func main() {
	q, r := div(13, 3)
	fmt.Println(q, r)
	fmt.Println(div(13, 3))
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
	fmt.Println(sum(1, 2))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
