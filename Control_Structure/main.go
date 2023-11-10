package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	i := 0
	for ; i < 50; i++ {
		fmt.Println(z)
		// create loss function
		z -= (z*z - x) / (2 * z)
		// break if squareroot found
		if z - math.Sqrt(x) == 0 {
			break
		}
	}
	return z
}

func grade(marks int) {
	switch gmarks := marks / 10; gmarks {
	case 10:
		fmt.Println("Your grade is A+")
	case 9:
		fmt.Println("Your grade is A+")
	case 8:
		fmt.Println("Your grade is A")
	case 7:
		fmt.Println("Your grade is B+")
	case 6:
		fmt.Println("Your grade is B")
	case 5:
		fmt.Println("Your grade is D+")
	case 4:
		fmt.Println("Your grade is D")
	default:
		fmt.Println("You failed")	
	}
}

func main() {
	// calculating grade
	grade(40)
	// calculate square root without builtin function
	fmt.Println(Sqrt(152587890625))
}
