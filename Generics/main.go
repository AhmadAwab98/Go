package main

import "fmt"

// contains checks if a value is present in a slice 
func contains[T comparable](list []T, value T) string {
	for _, item := range list {
		if item == value {
			return "Yes"
		}
	}
	return "No"
}

func main() {
	intList := []int{1, 3, 0, 9, 1, 7, 8, 5}
	strList := []string{"awab", "ali", "saad", "asad", "haris"}
	fltList := []float64{1.2, 2.97, 3.01, 2.07, 2.24}

	// Check if a specific value is present in the slices
	fmt.Println("Does Int List Contains 3?", contains(intList, 3))
	fmt.Println("Does Int List Contains 6?", contains(intList, 6)) 
	fmt.Println("Does String List Contains ali?", contains(strList, "ali"))
	fmt.Println("Does String List Contains aqib?", contains(strList, "aqib"))
	fmt.Println("Does Float List Contains 3.01?", contains(fltList, 3.01))
	fmt.Println("Does Float List Contains 2.57?", contains(fltList, 2.57))
}
