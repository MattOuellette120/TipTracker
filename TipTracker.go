package main

import (
	"fmt"
	"strconv"
)

var mileage float64 = 0.00
var totalTips float64 = 0.00
var tipString string
func main () {
	fmt.Println("What is the mileage rate?")
	fmt.Scan(&mileage)
	fmt.Println(fmt.Sprintf("Mileage is $%.2f", mileage))
	fmt.Println("Enter \"exit\" to end shift")
	for numOfDel := 1.0 ; ; numOfDel++ {
		fmt.Println(fmt.Sprintf("Enter the tip for delivery #%.0f", numOfDel))
		fmt.Scan(&tipString)
		if tipString == "exit" {
			fmt.Println("End of shift")
			fmt.Println(fmt.Sprintf("Number of deliveries: %.0f, Total Tips: $%.2f, Total Earnings: $%.2f", numOfDel, totalTips, totalTips+(mileage*numOfDel)))
			break
		} else if tip, err := strconv.ParseFloat(tipString, 64); err == nil {
			totalTips += tip
			fmt.Println(fmt.Sprintf("Total Tips: $%.2f", totalTips))
			
		}


	}
}
