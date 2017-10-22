package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Variable Declaration
var inputString string
var helpString = "\nCommand list:\nhelp		Displays this help menu\nexit		Ends the shift and displays the final totals\nlist		Lists all current orders with ID and tip amount\ncancel		Allows user to cancel an order with its order ID. Cancelling an order sets the tip amount to $0\n"		//Help Menu


func main () {
	orders := make([]Order, 0, 1)		//Slice of all orders for the shift
	
	fmt.Println("\"Mileage\" refers to the amount of money, per delivery, a driver is given for gas compensation.") 
	var mileage float64	//Declared mileage outside the loop so it can be used in the final calculation
	for {		//Mileage input loop start
		fmt.Println("How much is the current mileage?")
		fmt.Scan(&inputString)
		cleanInputString := strings.TrimPrefix(inputString, "$")		//In case the user adds the $ to their input, trims $
		m, err := strconv.ParseFloat(cleanInputString, 64)
		mileage = m		//Assigning the ParseFloat directly to mileage shadows the variable outside the loop
		if err == nil {
			break
		} else {
			fmt.Println("\nMileage must be in the form of a float. Mileage is usually between $1.20 and $1.30")	//Error handling
		}
	}		//Mileage input loop end
	fmt.Println("\nEnter \"help\" for command list, enter \"exit\" to end shift.")
	
	for {
		fmt.Println("Enter the tip for delivery #", strconv.Itoa(len(orders)+1))
		fmt.Scan(&inputString)
		cleanInputString := strings.TrimPrefix(inputString, "$")		//Same with mileage input, trims off potential $'s
		switch {		//Interpret different input commands
			case inputString == "exit":		//End shift and print totals
				clockOut(orders, mileage)
				break
			case inputString == "help":		//Display help menu
				fmt.Println(helpString)
				break
			case inputString == "list":		//Lists everything
				listAll(mileage, orders)
				break
			case inputString == "cancel":		//Cancel an order using order's ID
				fmt.Println("\nHere's a list of all deliveries made. Please input the order number of the delivery you want cancelled.")
				listOrders(orders)
				fmt.Scan(&inputString)
				id, err := strconv.Atoi(inputString)
				if err == nil {
					cancelOrder(orders, id-1)
					fmt.Println(fmt.Sprintf("Order has been cancelled\nNew Total Tips: $%.2f\n", getTotalTips(orders)))
				} else {
					fmt.Println("Invalid Input\n")
				}
				break
			default:		//Append new order to orders slice setting tip and id
				tip, err := strconv.ParseFloat(cleanInputString, 64)
				if err == nil {
					id := len(orders)+1
					newOrder := Order{tip, id}
					orders = append(orders, newOrder)
					fmt.Println(fmt.Sprintf("\nTotal Tips: $%.2f", getTotalTips(orders)))		//Handy display of current total
				} else {
					fmt.Println("Invalid Input\n")
				}
		}		//End of switch case
		if inputString == "exit" {
			break
		}
	}		//End of tip input For Loop
}		//End of main func


type Order struct {		//Order struct
	tip float64
	id int
}

func (o Order) getTip() float64 {		//return tip
	return o.tip
}

func cancelOrder(o []Order, id int) {		//Takes in the orders slice and id, sets tip to 0
	o[id].tip = 0.0
	return
}

func clockOut (o []Order, m float64) {		//Clock Out function
		t := getTotalTips(o)
		fmt.Println(fmt.Sprintf("\nShift Complete!\nTotal Deliveries: %d, Total Tips: $%.2f, Total Earning: $%.2f", len(o), t, t+(float64(len(o))*m)))
	return
}

func getTotalTips (o []Order) float64{		//Returns total of tips
	totalTips := 0.00
	for i := 0; i < len(o); i++ {
		totalTips += o[i].getTip()
	}
	return totalTips

}

func listAll(m float64, o []Order) {		//List all orders and mileage
	fmt.Println(fmt.Sprintf("\nMileage is $%.2f", m))
	listOrders(o)
	return
}

func listOrders(o []Order) {		//Lists just orders
	if len(o) == 0 {
		fmt.Println("You haven't taken any deliveries yet!\n")
	} else {
		for i := 0; i < len(o); i++ {
			fmt.Println(fmt.Sprintf("Order #%d: $%.2f", o[i].id, o[i].tip))
		}
	}
	return
}
