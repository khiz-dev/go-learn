package exercises

import (
	"io"
	"log"
	"os"
)

func fileLen(file string) int {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	data := make([]byte, 2048)
	defer f.Close()
	t := 0
	for {
		count, err := f.Read(data)
		t += count
		if err != nil {
			if err != io.EOF {
				log.Fatal("Error: ", err)
			}
			break
		}
	}
	return t

}

func Exercise2() {
	println("Exercise 2 ----")
	println(fileLen("exercises/generated.txt"))
}
