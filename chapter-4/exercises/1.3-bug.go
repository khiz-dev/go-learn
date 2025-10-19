package exercises

import "fmt"

func Exercise3() {
	fmt.Println("Exercise 3 ---- ")
	var total int
	for i := 0; i < 10; i++ {
		total = total + 1
		fmt.Println("Total: ", total)
	}
}
