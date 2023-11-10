package main

import (
	"fmt"
)

func main() {
	// array of revenue data
	revenueInDollars := []int{1800, 1500, 1800, 2000, 1600, 1900, 3100, 3000, 1400, 1500, 2000, 3200}

	// calculate average revenue
	totalrevenue := 0
	for _, v := range revenueInDollars {
		totalrevenue += v
	}
	avgrevenue := totalrevenue / len(revenueInDollars)

	// months with more than average revenue
	var aboveAvgMonths []int
	for i, v := range revenueInDollars {
		if v > avgrevenue {
			aboveAvgMonths = append(aboveAvgMonths, i+1)
		}
	}

	fmt.Println("Average monthly revenue in Dollars: \n", float64(avgrevenue))
	fmt.Println("Months with more than average revenue: \n", aboveAvgMonths)
}