package main

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) (r int) {
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && roman[string(s[i])] < roman[string(s[i+1])] {
			r += roman[string(s[i+1])] - roman[string(s[i])]
			i++
		} else {
			r += roman[string(s[i])]
		}
	}

	return r
}
