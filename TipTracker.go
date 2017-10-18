package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"io"
)

var mileage float64 = 0.00
var totalTips float64 = 0.00

func main () {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What is the mileage rate?")
	fmt.Scan(&mileage)
	fmt.Println(fmt.Sprintf("Mileage is $%.2f", mileage))
	fmt.Println("Enter an empty string to end shift")
	for numOfDel := 1.0 ; ; numOfDel++ {
		fmt.Println(fmt.Sprintf("Enter the tip for delivery #%.0f", numOfDel))
		tipString := scanner.Text()
		if len(tipString) == 0 {
			break
		}
		if err := scanner.Err(); err != nil {
			if err != io.EOF {
				fmt.Fprintln(os.Stderr, err)
			}
		}
		if tip, err := strconv.ParseFloat(tipString, 64); err == nil {
			totalTips += tip
		}
		fmt.Println(fmt.Sprintf("Total Tips: $%.2f", totalTips))
		if numOfDel == 10 {
			fmt.Println("End of shift")
			fmt.Println(fmt.Sprintf("Number of deliveries: %.0f, Total Tips: $%.2f, Total Earnings: $%.2f", numOfDel, totalTips, totalTips+(mileage*numOfDel)))
			break
		}


	}
}
