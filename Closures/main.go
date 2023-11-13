package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// initialize f1 and f2 
	f1, f2 := 0, 0
	return func() int {
		// change values according to fibonacci series
		f1, f2 = f2, f1 + f2
		if f2 == 0 {
			f2 = 1
		}
		return f1 
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
