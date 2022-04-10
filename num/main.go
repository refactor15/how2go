package main

import (
	"fmt"
	"math"
	"math/rand"
	"time")

func test(j, k int) int {
	if j == k {
		return 0
	} else {
		c := math.Abs(float64(j - k))
		if c < 10 {
			return int(math.Pow(c, 2))
		} else {
			d := float64(int(c/10) + int(math.Mod(c, 10)))
			if d < 10 {
				return int(math.Pow(d, 2))
			} else {
				e := float64(int(d/10) + int(math.Mod(d, 10)))
				return int(math.Pow(e, 2))
			}
		}
	}
}

func main() {
	var isGameover string
	for isGameover != "n" {
		var guess int
		var n int = 6
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(100)
		fmt.Printf("Guess a number between 0 and 99,you have %d chances.\n", n)
		defer fmt.Println("The answer is:", randomNumber)
		for gCount := 0; gCount < n; gCount++ {
		    fmt.Printf("No.%d guess:", gCount+1)
			fmt.Scanf("%d", &guess)
			if test(guess, randomNumber) == 0 {
				fmt.Printf("You got it in %d times!\n", gCount+1)
				break
			} else {
				fmt.Println("Try again.", test(guess, randomNumber))
			}
		}
		fmt.Println("do you want to play again? (y/n)")
		fmt.Scanf("%s", &isGameover)
	}
}
