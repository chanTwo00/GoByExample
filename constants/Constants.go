// Go supports constants of character, string, boolean, and numeric values.
// const declares a constant value.
// A const statement can appear anywhere a var statement can.
// Constant expressions perform arithmetic with arbitrary precision.
// A numeric constant has no type until it’s given one, such as by an explicit conversion.

package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n 
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

