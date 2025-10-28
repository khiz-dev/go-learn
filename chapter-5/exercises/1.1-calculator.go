package exercises

import (
	"log"
	"strconv"
)

func add(x int, y int) int      { return x + y }
func subtract(x int, y int) int { return x - y }
func multiply(x int, y int) int { return x * y }
func divide(x int, y int) int   { return x / y }

func Exercise1() {
	// This is just a placeholder to indicate where the exercise code would go.
	// The actual implementation of the calculator functions is above.

	calculator := map[string]func(int, int) int{
		"+": add,
		"-": subtract,
		"*": multiply,
		"/": divide,
	}

	calculations := [][]string{
		{"10", "+", "5"},
		{"10", "-", "5"},
		{"10", "*", "5"},
		{"10", "/", "5"},
		{"10", "/", "0"},
	}

	for _, i := range calculations {
		x, err := strconv.Atoi(i[0])
		if err != nil {
			log.Fatal("Invalid first expression")
		}
		y, err := strconv.Atoi(i[2])
		if err != nil {
			log.Fatal("Invalid second expression")
		}
		op := i[1]

		if operation, exists := calculator[op]; exists {
			if op == "/" && y == 0 {
				println("Error: Division by zero")
			} else {
				result := operation(x, y)
				println(x, op, y, "=", result)
			}
		} else {
			println("Error: Unknown operation", op)
		}
	}
}
