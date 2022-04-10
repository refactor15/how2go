package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func (p point) ToString() string {
	return fmt.Sprintf("x = %d; y = %d", p.x, p.y)
}

func (p *point) Move(right, down int) {
	p.x += right
	p.y += down
}

var p = point{1, 2}
var str1 = p.ToString()

func main() {
	fmt.Println(str1)
	(&p).Move(5, -1)
	var str2 = p.ToString()
	fmt.Println(str2)
}
