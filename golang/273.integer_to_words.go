package main

import "strings"

var sub20 = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var tens = []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
var thousands = []string{"", "Thousand", "Million", "Billion"}

func helper(num int) string {
	if num == 0 {
		return ""
	}

	if num < 20 {
		return sub20[num] + " "
	} else if num < 100 {
		return tens[num/10] + " " + helper(num%10)
	} else {
		return sub20[num/100] + " Hundred " + helper(num%100)
	}
}

func numberToWords(num int) (r string) {

	if num == 0 {
		return "Zero"
	}

	for _, el := range thousands {
		if num%1000 != 0 {
			r = helper(num%1000) + el + " " + r
		}
		num = num / 1000
	}

	return strings.TrimSpace(r)
}
