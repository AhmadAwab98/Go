package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

// error returns the error message for ErrNegativeSqrt
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// return the square root and an error if the input is negative.
func Sqrt(x float64) (float64, error) {
	if x > 0 {
		z := 1.0
		i := 0
		for ; i < 50; i++ {
			// create loss function
			z -= (z*z - x) / (2 * z)
			// break if squareroot found
			if z - math.Sqrt(x) == 0 {
				break
			}
		}
		return float64(z), nil
	}
	// return an error for negative input
	return 0, ErrNegativeSqrt(x)
}

func main() {
	value, err := Sqrt(2)
	if err != nil {
		fmt.Println(err)
	}
	if value > 0	{
		fmt.Println(value)
	}
	value, err = Sqrt(-2)
	if err != nil {
		fmt.Println(err)
	}
	if value > 0 	{
		fmt.Println(value)
	}
}