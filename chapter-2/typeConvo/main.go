package main

import (
	"fmt"
	"math"
)

func main() {
	var todayBitcoinPrice float64 = 27345.23
	var todayBitcoinPriceInt int = int(todayBitcoinPrice)
	var todayBitcoinPriceRoundedUp float64 = math.Ceil(todayBitcoinPrice)

	fmt.Println("Today's Bitcoin Price: ", todayBitcoinPrice)
	fmt.Println("Today's Bitcoin Price (int): ", todayBitcoinPriceInt)
	fmt.Println("Today's Bitcoin Price (rounded up): ", todayBitcoinPriceRoundedUp)
}

/*
Today's Bitcoin Price:  27345.23
Today's Bitcoin Price (int):  27345
Today's Bitcoin Price (rounded up):  27346
*/

/*
The type conversion and rounding up are not the same.
*/
