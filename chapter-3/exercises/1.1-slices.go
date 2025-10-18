package exercises

import "fmt"

func Exercise1() {
	fmt.Println("Exercise 1 -- ")
	var greetings = []string{"Hello", "Hola", "नमस्कार", "你好", "Γειά σου"}
	fSlice := greetings[:2]
	sSlice := greetings[1:4]
	tSlice := greetings[3:]

	fmt.Println("First two greetings: ", fSlice)
	fmt.Println("Middle three greetings: ", sSlice)
	fmt.Println("Last greeting: ", tSlice)
}
