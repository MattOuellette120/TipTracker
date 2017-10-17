package main

import (
	"fmt"
	"reflect"
)

var mileage float64 = 0.00
var totalTips float64 = 0.00
var numOfDel float64 = 1.00
var shiftDone bool = false

func main () {
	fmt.Println("What is the mileage rate?")
	fmt.Scan(&mileage)
	fmt.Println(fmt.Sprintf("Mileage is $%.2f", mileage))
	fmt.Println("Input \"exit\" to end shift")
	for ; shiftDone == true ; numOfDel++ {
		fmt.Println(fmt.Sprintf("Enter the tip for delivery #%d", numOfDel))
		var tip 
		fmt.Scan(&tip)
		var tipType string = reflect.TypeOf(tip).String()
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
		}


	}
}
