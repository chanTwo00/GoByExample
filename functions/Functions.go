// Functions are central in Go. We’ll learn about functions with a few different examples.
// Here’s a function that takes two ints and returns their sum as an int.
// Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
// When you have multiple consecutive parameters of the same type, you may omit the type name for the 
// like-typed parameters up to the final parameter that declares the type.
// Call a function just as you’d expect, with name(args).

package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusplus(1, 2, 3)
	fmt.Println("1+2+3=", res)
}

