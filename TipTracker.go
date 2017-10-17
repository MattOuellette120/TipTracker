package main

import (
	"fmt"
	//"reflect"
)

var mileage float64 = 0.00
var totalTips float64 = 0.00
var tip 

func main () {
	fmt.Println("What is the mileage rate?")
	fmt.Scan(&mileage)
	fmt.Println(fmt.Sprintf("Mileage is $%.2f", mileage))
	fmt.Println("Input \"exit\" to end shift")
	for numOfDel := 1.0 ; ; numOfDel++ {
		fmt.Println(fmt.Sprintf("Enter the tip for delivery #%.0f", numOfDel))
		fmt.Scan(&tip)
/*		var tipType string = reflect.TypeOf(tip).String()
		if tipType == "string" {
			if tip == "exit" {
				shiftDone = true
				fmt.Println("End of shift")
				fmt.Println(fmt.Sprintf("Number of deliveries: %d, Total Tips: $%.2f, Total Earnings: $%.2f", numOfDel, totalTips, totalTips+(mileage*numOfDel)))
			} else {
				fmt.Println("Invalid string")
				numOfDel--
			}
		} else if  tipType == "float64" {
				totalTips += tip
		} else if tipeType == "int" {
			floatTip := float64(tip)
			totalTips += floatTip
		} else {
			fmt.Println("Invalid Input")
			numOfDel--
		}*/
		if tip == "exit" {
			break
		}
		totalTips += tip
		fmt.Println(fmt.Sprintf("Total Tips: $%.2f", totalTips))
		if numOfDel == 10 {
			fmt.Println("End of shift")
			fmt.Println(fmt.Sprintf("Number of deliveries: %.0f, Total Tips: $%.2f, Total Earnings: $%.2f", numOfDel, totalTips, totalTips+(mileage*numOfDel)))
			break
		}


	}
}
