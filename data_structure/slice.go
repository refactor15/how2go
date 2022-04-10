package main
import "fmt"
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func main() {
	var a []int
	printSlice(a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice(a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice(a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice(a)
}
