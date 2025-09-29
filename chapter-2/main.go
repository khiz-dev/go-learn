package main

import (
	"fmt"
)

func printMoney(money int) {
	fmt.Println("Money: ", formatMoney(money))
}

func formatMoney(money int) string {
	return fmt.Sprintf("£%.2f", float32(money)/100)
}

func bankTransactionWithFloat() {
	var money float32 = 100.50
	fmt.Println("Money: ", money)
	money += 12.44
	fmt.Println("£12.44 is deposited into your account. Total balance is now £", money)
	money -= 20.00
	fmt.Println("$20.00 is withdrawn from your account. Total balance is now $", money)
	var interest float32 = money * 0.05
	money += interest
	fmt.Println("5% interest (", interest, ") is added to your account. Total balance is now $", money)
	fmt.Println("Money: ", money)
}

func bankTransactionWithInt() {
	var money int = 10050 // representing £100.50 in pence
	printMoney(money)
	money += 1244
	fmt.Println("£12.44 is deposited into your account. Total balance is now £", formatMoney(money))
	money -= 2000
	fmt.Println("£20.00 is withdrawn from your account. Total balance is now £", formatMoney(money))
	var interest int = money * 5 / 100
	money += interest
	fmt.Println("5% interest (", formatMoney(interest), ") is added to your account. Total balance is now £", formatMoney(money))
	printMoney(money)
}

func main() {
	fmt.Println("--- 15th September 2024 ---")
	bankTransactionWithFloat()
	fmt.Println("--- 16th September 2024 ---")
	bankTransactionWithInt()
}

/* Output of the program:
--- 15th September 2024 ---
Money:  100.5
£12.44 is deposited into your account. Total balance is now £ 112.94
$20.00 is withdrawn from your account. Total balance is now $ 92.94
5% interest ( 4.6470003 ) is added to your account. Total balance is now $ 97.587006
Money:  97.587006
--- 16th September 2024 ---
Money:  £100.50
£12.44 is deposited into your account. Total balance is now £ £112.94
£20.00 is withdrawn from your account. Total balance is now £ £92.94
5% interest ( £4.64 ) is added to your account. Total balance is now £ £97.58
Money:  £97.58
*/
