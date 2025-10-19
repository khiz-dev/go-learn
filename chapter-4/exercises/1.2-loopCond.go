package exercises

import (
	"fmt"
	"math/rand"
)

func Exercise2() {
	fmt.Println("Exercise 2 ---- ")
	randInts := make([]int, 0, 100)
	for i := 0; i <= 100; i++ {
		randInt := rand.Intn(100)
		switch {
		case (randInt%2 == 0) && (randInt%3 == 0):
			fmt.Println("Divisble by both 2 and 3, it must be a six: ", randInt)
		case randInt%2 == 0:
			fmt.Println("Divisble by 2")
		case randInt%3 == 0:
			fmt.Println("Divisible by 3")
		default:
			fmt.Println("Never mind")
		}
		randInts = append(randInts, rand.Intn(100))
	}
	fmt.Println("Rand Slice of size 100: ", randInts)
}
