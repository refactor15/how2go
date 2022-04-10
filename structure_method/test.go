package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func main() {
	fmt.Printf("gg %g y.\n", math.Sqrt(2))
	fmt.Println(math.Pi)
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(split(17))
	fmt.Println(add(2, 3))
}
