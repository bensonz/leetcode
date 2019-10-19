package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	profit, boughtPrice := 0, 0
	bought := false
	prices = append(prices, 0) // add a 0 at the end
	for i := 0; i < len(prices)-1; i++ {
		if bought {
			// find a sell date
			if prices[i+1] < prices[i] {
				// sell at day i
				fmt.Println("Sold at ", prices[i])
				profit += prices[i] - boughtPrice
				bought = false
			}
		} else {
			// find a buy date
			if prices[i+1] > prices[i] {
				fmt.Println("bought at ", prices[i])
				boughtPrice = prices[i]
				bought = true
			}
		}
	}
	return profit
}

func main() {
	p := []int{1, 2, 3, 4, 0, 0, 5}
	fmt.Println(maxProfit(p))
}
