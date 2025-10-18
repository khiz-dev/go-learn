package exercises

import (
	"fmt"
	"strings"
)

func Exercise4() {
	fmt.Println("Exercise 4 -- ")
	letterOcc := map[string]int{}
	wordOcc := map[string]int{}

	fmt.Println("Initial Occurrences Map: ", letterOcc)
	fmt.Println("Initial Words Occurrences Map: ", wordOcc)

	sentence := "K8s Operator Development is a great way to learn about both the GO programming language and Kubernetes internals. K8s Operators help automate complex tasks in Kubernetes."

	// Characters count
	for _, char := range sentence {
		letterOcc[strings.ToLower((string(char)))]++
	}

	for _, word := range strings.Split(sentence, " ") {
		wordOcc[strings.ToLower(word)]++
	}

	//

	fmt.Println("Letter occurrences Map: ", letterOcc)
	fmt.Println("Word occurrences map: ", wordOcc)
	fmt.Println("")
}
