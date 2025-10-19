package exercises

import "fmt"

func Exercise0() {
	fmt.Println("Exercise 0 -----")

outer:
	for i := 0; i <= 20; i += 1 {
		switch j := i % 5; j {
		case 0:
			fmt.Println("Multiple of 5: ", i)
			i++
		case 1, 2:
			fmt.Println("Within first two digits: ", i)
			i += 2
		case 3, 4:
			fmt.Println("Within the last two digits: ", i)
			break outer
		}
	}
}
