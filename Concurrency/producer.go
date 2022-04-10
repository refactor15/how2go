package main
import (
	"fmt"
	"time"
)
func Producer(factor int,out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * factor
	}
}
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
func main() {
	ch := make(chan int,8)
	go Producer(-1,ch)
	go Producer(1,ch)
	go Consumer(ch)
	time.Sleep(100 * time.Microsecond)
}

