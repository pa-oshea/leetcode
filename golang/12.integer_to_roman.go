package main

var romanMap = []map[string]int{
	{"M": 1000},
	{"CM": 900},
	{"D": 500},
	{"CD": 400},
	{"C": 100},
	{"XC": 90},
	{"L": 50},
	{"XL": 40},
	{"X": 10},
	{"IX": 9},
	{"V": 5},
	{"IV": 4},
	{"I": 1},
}

func intToRoman(num int) (r string) {
	for _, value := range romanMap {
		for key, el := range value {
			for num >= el {
				r += key
				num -= el
			}
		}

	}

	return
}
