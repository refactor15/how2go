package main
import "fmt"

func main(){
q := []int{2, 3, 5, 7, 11, 13}
r := []bool{true, false, true, true, false, true}
t := []struct {
	i int
	b bool
}{
	{2, true},
	{3, false},
	{5, true},
	{7, true},
	{11, false},
	{13, true},
}	
		var s []int
		fmt.Println(s, len(s), cap(s))
		if s == nil {
			fmt.Println("nil!")
		}
		fmt.Println(q,r,t)
}
