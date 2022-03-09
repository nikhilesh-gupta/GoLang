package main

import "fmt"

func factorialViaChannel(n int, channelVariable chan func() (int, int)) {
	val := 1
	if n == 0 {
		channelVariable <- (func() (int, int) { return n, val })
	} else {
		for i := 1; i <= n; i++ {
			val *= i
		}
		channelVariable <- (func() (int, int) { return n, val })
	}
}

func main() {
	var factorial int
	channelVariable := make(chan func() (int, int), 1)

	fmt.Print(">> ")
	fmt.Scanln(&factorial)

	go factorialViaChannel(factorial, channelVariable)
	val, res := (<-channelVariable)()
	fmt.Printf("Factorial of %d is %d\n", val, res)
}
