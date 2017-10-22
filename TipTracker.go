package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Variable Declaration
var mileageString string
var tipString string
var helpString = "Command list:\nhelp		Displays this help menu\nexit		Ends the shift and displays the final totals"


func main () {
	orders := make([]Order, 0, 1)
	
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
	fmt.Println("Enter \"help\" for command list, enter \"exit\" to end shift.")
	
	for {
		fmt.Println("Enter the tip for delivery #", len(orders)+1)
		fmt.Scan(&tipString)
		cleanTipString := strings.TrimPrefix(tipString, "$")		//Same with mileage input, trims off potential $'s
		switch {
			case tipString == "exit":
				clockOut(orders, mileage)
				break
			case tipString == "help":
				fmt.Println(helpString)
				break
			default:
				tip, err := strconv.ParseFloat(cleanTipString, 64)
				if err == nil {
					newOrder := Order{tip}
					orders = append(orders, newOrder)
					fmt.Println(fmt.Sprintf("Total Tips: $%.2f", getTotalTips(orders)))		//Handy display of current total
				} else {
					fmt.Println("Invalid Input")
				}
		}		//End of switch case
		if tipString == "exit" {
			break
		}
	}		//End of tip input For Loop
}		//End of main func


type Order struct {		//Order struct
	tip float64
}

func (o Order) getTip() float64 {
	return o.tip
}

func (o Order) cancelOrder() {		//To be implemented
	o.tip = 0.0
	return
}

func clockOut (o []Order, m float64) {		//Clock Out function
		t := getTotalTips(o)
		fmt.Println(fmt.Sprintf("Shift Complete!\nTotal Deliveries: %d, Total Tips: $%.2f, Total Earning: $%.2f", len(o), t, t+(float64(len(o))*m)))
	return
}

func getTotalTips (o []Order) float64{		//Returns total of tips
	totalTips := 0.00
	for i := 0; i < len(o); i++ {
		totalTips += o[i].getTip()
	}
	return totalTips

}
