package exercises

import "fmt"

func Exercise2() {
	fmt.Println("Exercise 2 -- ")
	const bankLimit = 10
	var currentBankLimit = bankLimit
	var savingBankLimit = bankLimit * 0.86

	fmt.Println("Current Bank Limit: ", currentBankLimit)
	fmt.Println("Saving Bank Limit: ", savingBankLimit)
}
