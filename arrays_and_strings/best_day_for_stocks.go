/*
121. Best Time to Buy and Sell Stock

You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing
a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/

package main

import "math"

// Memory: O(n)
// Space: O(n)
func maxProfit(prices []int) int {
	var min, profit, maxProfit int = math.MaxInt, 0, 0
	for _, p := range prices {
		if p < min {
			min = p // update min price if necessary
		}
		profit = p - min // update current profit
		if profit > maxProfit {
			maxProfit = profit // update maxProfit if necessary
		}
	}
	return maxProfit
}
