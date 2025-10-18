package exercises

import "fmt"

func Exercise2() {
	fmt.Println("Exercise 2 -- Slices/Runes")
	var message = "😄 and 😄"
	var forthRune = []rune(message)[3]
	fmt.Println("Message: ", message)
	fmt.Println("The forth character in the message is: ", message[3])
	fmt.Printf("The fourth rune in the message is: %c\n", forthRune)
	fmt.Println("")
}
