package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Variable Declaration
var mileageString string
var totalTips float64 = 0.00
var tipString string

func main () {
	fmt.Println("\"Mileage\" refers to the amount of money, per delivery, a driver is given for gas compensation.") 
	var mileage float64	//Declared mileage outside the loop so it can be used in the final calculation
	for {		//Mileage input loop start
		fmt.Println("How much is the current mileage?")
		fmt.Scan(&mileageString)
		cleanMileageString := strings.TrimPrefix(mileageString, "$")		//In case the user adds the $ to their input, trims $
		m, err := strconv.ParseFloat(cleanMileageString, 64)
		mileage = m		//Assigning the ParseFloat directly to mileage shadows the variable outside the loop
		if err == nil {
			break
		} else {
			fmt.Println("Mileage must be in the form of a float. Mileage is usually between $1.20 and $1.30")	//Error handling
		}
	}		//Mileage input loop end
	fmt.Println(fmt.Sprintf("Mileage is $%.2f", mileage))
	fmt.Println("Enter \"exit\" to end shift")
	
	for numOfDel := 1.0 ; ; numOfDel++ {
		fmt.Println(fmt.Sprintf("Enter the tip for delivery #%.0f", numOfDel))
		fmt.Scan(&tipString)
		cleanTipString := strings.TrimPrefix(tipString, "$")		//Same with mileage input, trims off potential $'s
		if tipString == "exit" {		//If "exit" was entered, ends the loop and prints final totals
			numOfDel--		//Minus one, because it would use the current loop's number otherwise
			fmt.Println("End of shift")
			fmt.Println(fmt.Sprintf("Number of deliveries: %.0f, Total Tips: $%.2f, Total Mileage Earned: $%.2f, Total Earnings: $%.2f", numOfDel, totalTips, numOfDel*mileage, totalTips+(mileage*numOfDel)))
			break
		}
		tip, err := strconv.ParseFloat(cleanTipString, 64)
		if err == nil {		//If no errors parsing in the float, adds the tip input to the running total
			totalTips += tip
			fmt.Println(fmt.Sprintf("Total Tips: $%.2f", totalTips))		//Handy display of current total
		} else {
			fmt.Println("Invalid Input")
			numOfDel--		//If invalid input, steps back the loop
		}


	}
}



