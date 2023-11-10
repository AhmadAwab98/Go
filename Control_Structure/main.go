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
	// grade is A+ if marks in between 90 and 100
	case 10:
		fmt.Println("Your grade is A+")
	case 9:
		fmt.Println("Your grade is A+")
	// grade is A if marks in between 80 and 89
	case 8:
		fmt.Println("Your grade is A")
	// grade is B+ if marks in between 70 and 79
	case 7:
		fmt.Println("Your grade is B+")
	// grade is B if marks in between 60 and 69
	case 6:
		fmt.Println("Your grade is B")
	// grade is C if marks in between 50 and 59
	case 5:
		fmt.Println("Your grade is C")
	// grade is D if marks in between 40 and 49
	case 4:
		fmt.Println("Your grade is D")
	// fail if marks less than 40
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
