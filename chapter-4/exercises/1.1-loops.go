package exercises

import (
	"fmt"
	"math/rand"
)

func Exercise1() {
	fmt.Println("Exercise 1 ---- ")
	randInts := make([]int, 0, 100)
	for i := 0; i <= 100; i++ {
		randInts = append(randInts, rand.Intn(100))
	}
	fmt.Println("Rand Slice of size 100: ", randInts)
}
