package main

import (
	"fmt"
	"strconv"
	"strings"
)

var mileageString string
var totalTips float64 = 0.00
var tipString string

func main () {
	fmt.Println("\"Mileage\" refers to the amount of money, per delivery, a driver is given for gas compensation.")
	var mileage float64
	for {
		fmt.Println("How much is the current mileage?")
		fmt.Scan(&mileageString)
		cleanMileageString := strings.TrimPrefix(mileageString, "$")
		m, err := strconv.ParseFloat(cleanMileageString, 64)
		mileage = m
		if err == nil {
			break
		} else {
			fmt.Println("Mileage must be in the form of a float. Mileage is usually between $1.20 and $1.30")
		}
	}
	fmt.Println(fmt.Sprintf("Mileage is $%.2f", mileage))
	fmt.Println("Enter \"exit\" to end shift")
	
	for numOfDel := 1.0 ; ; numOfDel++ {
		fmt.Println(fmt.Sprintf("Enter the tip for delivery #%.0f", numOfDel))
		fmt.Scan(&tipString)
		cleanTipString := strings.TrimPrefix(tipString, "$")
		if tipString == "exit" {
			numOfDel--
			fmt.Println("End of shift")
			fmt.Println(fmt.Sprintf("Number of deliveries: %.0f, Total Tips: $%.2f, Total Mileage Earned: $%.2f, Total Earnings: $%.2f", numOfDel, totalTips, numOfDel*mileage, totalTips+(mileage*numOfDel)))
			break
		}
		tip, err := strconv.ParseFloat(cleanTipString, 64)
		if err == nil {
			totalTips += tip
			fmt.Println(fmt.Sprintf("Total Tips: $%.2f", totalTips))
		} else {
			fmt.Println("Invalid Input")
			numOfDel--
		}


	}
}



