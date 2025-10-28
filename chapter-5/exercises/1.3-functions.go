package exercises

func prefixer(pre string) func(s string) string {
	return func(s string) string {
		return pre + " " + s
	}
}

func Exercise3() {
	println("Exercise 3 ---- ")
	helloPrefix := prefixer("Hello")
	println(helloPrefix("Bob"))
	println(helloPrefix("Dylan"))
}
