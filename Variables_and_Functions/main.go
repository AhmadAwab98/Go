package main

import (
	"fmt"
	"math"
)

func isubtract_two(a, b int) int {

    return a - b
}

func iadd_three(a, b, c int) int {
    return a + b + c
}

func pythagoras_theorem(a, b int) (c float64) {
	// convert a and b to float64
	d, e := float64(a), float64(b)
	//computing c
	c := math.Sqrt(d*d + e*e)
	return
}

func main() {
	x, y, z := 27, 9, 18
	// subtract two int type variables 
	subtraction_int := isubtract_two(x, y)
	fmt.Println("27 - 9 =", subtraction_int)
	// add three int type variables
	addition_int := iadd_three(x, y, z)
	fmt.Println("27 + 9 + 18 =", addition_int)
	
	// compute hypotenuse using pythagoras theorem
	c := pythagoras_theorem(x,y)
	fmt.Println("Hypotenuse of given triangle is:", c)
}