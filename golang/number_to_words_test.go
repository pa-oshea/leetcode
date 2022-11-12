package main

import "testing"

func TestNumberToWords(t *testing.T) {

	assertEqual(t, numberToWords(10), "Ten", "")
	assertEqual(t, numberToWords(123), "One Hundred Twenty Three", "")
	assertEqual(t, numberToWords(1000), "One Thousand", "")
	assertEqual(t, numberToWords(10000), "Ten Thousand", "")
	assertEqual(t, numberToWords(100000), "One Hundred Thousand", "")
	assertEqual(t, numberToWords(1000000), "One Million", "")
	assertEqual(t, numberToWords(12345), "Twelve Thousand Three Hundred Forty Five", "")
	assertEqual(t, numberToWords(50868), "Fifty Thousand Eight Hundred Sixty Eight", "")
	assertEqual(t, numberToWords(58868), "Fifty Eight Thousand Eight Hundred Sixty Eight", "")
	assertEqual(t, numberToWords(1000000000), "One Billion", "")

}
