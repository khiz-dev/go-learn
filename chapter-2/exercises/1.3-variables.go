package exercises

import "fmt"

func Exercise3() {
	fmt.Println("Exercise 3 -- ")

	// Assign the max value for these types
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI int64 = 9223372036854775807
	fmt.Println("Byte: ", b)
	fmt.Println("Byte +1: ", b+1)
	fmt.Println("Small Int: ", smallI)
	fmt.Println("Small Int +1: ", smallI+1)
	fmt.Println("Big Int: ", bigI)
	fmt.Println("Big Int +1: ", bigI+1)
}
