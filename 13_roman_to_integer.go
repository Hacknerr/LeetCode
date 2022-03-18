package main

// https://leetcode.com/problems/roman-to-integer/

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	sum := 0
	greatest := 0
	for i := len(s) - 1; i >= 0; i-- {
		letter := s[i]
		num := roman[string(letter)]
		if num >= greatest {
			greatest = num
			sum += num
			continue
		}
		sum -= num
	}
	return sum
}
