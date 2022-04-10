package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	var i int
	v := "hello"
	a, b := swap("hello", "world")
	fmt.Println(a, b, i)
	fmt.Printf("v type %T\n", v)
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
